package server

import (
	"github.com/alexliesenfeld/health"
	healthcheckhandler "ps-cats-social/internal/healthcheck/handler"
)

func (s *Server) setupRouter() {
	v1 := s.router.PathPrefix("/v1").Subrouter().StrictSlash(true)
	v1.HandleFunc("/health", health.NewHandler(healthcheckhandler.HealthCheck())).Methods("GET")
	v1.HandleFunc("/user/register", s.baseHandler.RunAction(s.userHandler.Register)).Methods("POST")
	v1.HandleFunc("/user/login", s.baseHandler.RunAction(s.userHandler.Login)).Methods("POST")

	v1.HandleFunc("/cat", s.baseHandler.RunActionAuth(s.catHandler.CreateCat)).Methods("POST")
	v1.HandleFunc("/cat/{id:[1-9][0-9]*}", s.baseHandler.RunActionAuth(s.catHandler.DeleteCat)).Methods("DELETE")
	v1.HandleFunc("/cat/{id:[1-9][0-9]*}", s.baseHandler.RunActionAuth(s.catHandler.UpdateCat)).Methods("PUT")
	v1.HandleFunc("/cat", s.baseHandler.RunActionAuth(s.catHandler.GetCat)).Methods("GET")

	v1.HandleFunc("/cat/match", s.baseHandler.RunActionAuth(s.catMatchHandler.MatchCat)).Methods("POST")
	v1.HandleFunc("/cat/match", s.baseHandler.RunActionAuth(s.catMatchHandler.GetMatches)).Methods("GET")
	v1.HandleFunc("/cat/match/approve", s.baseHandler.RunActionAuth(s.catMatchHandler.ApproveRequest)).Methods("POST")
	v1.HandleFunc("/cat/match/reject", s.baseHandler.RunActionAuth(s.catMatchHandler.RejectRequest)).Methods("POST")
	v1.HandleFunc("/cat/match/{id:[1-9][0-9]*}", s.baseHandler.RunActionAuth(s.catMatchHandler.DeleteMatch)).Methods("DELETE")

}
