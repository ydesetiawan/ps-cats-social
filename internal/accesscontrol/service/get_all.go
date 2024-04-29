package service

import (
	"ps-cats-social/internal/accesscontrol/dto/response"
	"ps-cats-social/pkg/base/app"
	"ps-cats-social/pkg/helper"
	"ps-cats-social/pkg/middleware"
	"strings"
)

func (s *service) GetAll(ctx *app.Context) (*response.AccessControlResponse, error) {
	roles := extractRoles(strings.ReplaceAll(ctx.Request.Header.Get("X-Consumer-Roles"), " ", ""))

	accessControlResponse := response.AccessControlResponse{}

	for _, rule := range middleware.AccessControlRules {
		if helper.ContainsAll(roles, rule.Roles) {
			applyRules(&accessControlResponse, rule.AllowedMethod)
		}
	}

	return &accessControlResponse, nil
}

func extractRoles(roleHeader string) []string {
	return strings.Split(roleHeader, ",")
}

func applyRules(response *response.AccessControlResponse, allowedMethods map[string][]string) {
	for path, methods := range allowedMethods {
		if strings.Contains(path, "reminders") {
			response.Reminder = generateAllowedMethods(methods)
		}
	}
}

func generateAllowedMethods(methods []string) response.AllowedMethods {
	allowedMethods := response.AllowedMethods{}
	for _, method := range methods {
		switch method {
		case "GET":
			allowedMethods.Read = true
		case "POST":
			allowedMethods.Create = true
		case "PUT":
			allowedMethods.Update = true
		case "DELETE":
			allowedMethods.Delete = true
		}
	}
	return allowedMethods
}
