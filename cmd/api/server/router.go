package server

import (
	"github.com/alexliesenfeld/health"
	healthcheckhandler "ps-cats-social/internal/healthcheck/handler"
)

func (s *Server) setupRouter() {
	v1 := s.router.PathPrefix("/v1").Subrouter().StrictSlash(true)
	v1.HandleFunc("/health", health.NewHandler(healthcheckhandler.HealthCheck())).Methods("GET")
	v1.HandleFunc("/user/register", s.baseHandler.RunAction(s.userHandler.RegisterUserHandler)).Methods("POST")
	v1.HandleFunc("/user/login", s.baseHandler.RunAction(s.userHandler.Login)).Methods("POST")

	v1.HandleFunc("/cat", s.baseHandler.RunActionAuth(s.catHandler.CreateCat)).Methods("POST")
	v1.HandleFunc("/cat/{id}", s.baseHandler.RunActionAuth(s.catHandler.DeleteCat)).Methods("DELETE")
	v1.HandleFunc("/cat/{id}", s.baseHandler.RunActionAuth(s.catHandler.UpdateCat)).Methods("PUT")
	v1.HandleFunc("/cat", s.baseHandler.RunActionAuth(s.catHandler.GetCat)).Methods("GET")

}
