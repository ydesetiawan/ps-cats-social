package authz

import (
	"fmt"
	"net/http"
	"strings"
)

func ExtractUserRoles(r *http.Request) ([]string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return nil, fmt.Errorf("missing auth header")
	}

	rolesHeader := strings.ReplaceAll(r.Header.Get("X-Consumer-Roles"), " ", "")
	if rolesHeader == "" {
		return nil, fmt.Errorf("missing roles header")
	}

	return strings.Split(rolesHeader, ","), nil
}
