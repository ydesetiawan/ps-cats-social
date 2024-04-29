package middleware

import (
	"context"
	"fmt"
	"golang.org/x/exp/slog"
	"net/http"
	"ps-cats-social/pkg/helper"
	"ps-cats-social/pkg/httphelper"
	"ps-cats-social/pkg/httphelper/response"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"ps-cats-social/pkg/authz"
)

type AccessControlRule struct {
	Roles         []string
	AllowedMethod map[string][]string
}

var AccessControlRules = map[string]AccessControlRule{
	"owner": {
		Roles: []string{"owner"},
		AllowedMethod: map[string][]string{
			"/public/lo/v1/reminders": {"GET", "POST", "PUT", "DELETE"},
		},
	},
	"ultimate": {
		Roles: []string{"ultimate"},
		AllowedMethod: map[string][]string{
			"/public/lo/v1/access_control_list": {"GET", "POST"},
			"/public/lo/v1/reminders":           {"GET", "POST"},
		},
	},
	"ultimateLm": {
		Roles: []string{"ultimate", "list_manager"},
		AllowedMethod: map[string][]string{
			"/public/lo/v1/access_control_list": {"GET", "POST", "PUT", "DELETE"},
			"/public/lo/v1/reminders":           {"GET", "POST", "PUT", "DELETE"},
		},
	},
}

func JWTAuthMiddleware(fn http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		jwtToken, err := extractJWTTokenFromHeader(r)
		if err != nil {
			slog.Error("Failed to extract JWT token from header", "error", err)
			writeUnauthorized(rw)
			return
		}

		claims, err := parseJWTToClaims(jwtToken)
		if err != nil {
			slog.Error("Failed to parse JWT token", "error", err)
			writeUnauthorized(rw)
			return
		}

		companyId, cidOk := claims["company_id"].(float64)
		userId, uidOk := claims["user_id"].(float64)
		if !cidOk || !uidOk {
			slog.Error("Invalid claims")
			writeUnauthorized(rw)
			return
		}

		user, err := constructUserInfo(r, companyId, userId)
		if err != nil {
			slog.Error("Failed to construct user info", "error", err)
			writeUnauthorized(rw)
			return
		}

		validRole, err := validateUserRoles(r.URL.Path, r.Method, user)
		if err != nil || !validRole {
			slog.Error("Invalid roles", "error", err)
			writeUnauthorized(rw)
			return
		}

		r2 := r.WithContext(context.WithValue(r.Context(), "user_info", user))
		slog.Debug("AUTHORIZED", "user_info", r2.Context().Value("user_info"))

		fn(rw, r2)
	}
}

func validateUserRoles(apiEndpoint string, method string, user map[string]interface{}) (bool, error) {
	roles, ok := user["roles"].([]string)
	if !ok || len(roles) == 0 {
		return false, fmt.Errorf("empty or invalid roles")
	}

	for _, rule := range AccessControlRules {
		if helper.ContainsAll(roles, rule.Roles) {
			for path, allowedMethods := range rule.AllowedMethod {
				if strings.HasPrefix(apiEndpoint, path) && sliceContainsString(allowedMethods, method) {
					return true, nil
				}
			}
		}
	}

	return false, nil
}

func sliceContainsString(slice []string, target string) bool {
	if slice == nil || len(slice) == 0 {
		return false
	}
	for _, val := range slice {
		if val == target {
			return true
		}
	}
	return false
}

func writeUnauthorized(rw http.ResponseWriter) {
	httphelper.WriteJSON(
		rw, http.StatusUnauthorized,
		response.WebResponse{
			Status:  http.StatusUnauthorized,
			Message: http.StatusText(http.StatusUnauthorized),
		},
	)
}

func extractJWTTokenFromHeader(r *http.Request) (string, error) {
	authToken := r.Header.Get("Authorization")
	if authToken == "" {
		return "", fmt.Errorf("missing auth token")
	}

	return authToken[len("Bearer "):], nil
}

func parseJWTToClaims(jwtToken string) (jwt.MapClaims, error) {
	token, _, err := jwt.NewParser().ParseUnverified(jwtToken, jwt.MapClaims{})
	if err != nil {
		return nil, err
	}

	// no need to verify 'token' signature since it already validated in authz kong plugin, just parse the token

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid jwt token")
	}

	return claims, nil
}

func constructUserInfo(r *http.Request, companyId, userId float64) (map[string]interface{}, error) {
	userRoles, err := authz.ExtractUserRoles(r)
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{
		"company_id": int(companyId),
		"user_id":    int(userId),
		"roles":      userRoles,
	}, nil
}
