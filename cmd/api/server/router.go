package server

import (
	"github.com/alexliesenfeld/health"
	"github.com/gorilla/mux"
	healthcheckhandler "ps-cats-social/internal/healthcheck/handler"
)

func (s *Server) setupRouter() {
	v1 := s.router.PathPrefix("/api/v1").Subrouter().StrictSlash(true)
	v1.HandleFunc("/health", health.NewHandler(healthcheckhandler.HealthCheck())).Methods("GET")

	publicLoApiV1 := s.router.PathPrefix("/public/lo/v1").Subrouter().StrictSlash(true)
	s.aclRoutes(publicLoApiV1)
}

func (s *Server) aclRoutes(publicLoApiV1 *mux.Router) {
	publicLoApiV1.HandleFunc("/access_control_list",
		s.baseHandler.RunActionAuth(s.accessControlListHandler.GetAccessControlListHandler)).Methods("GET")
}
