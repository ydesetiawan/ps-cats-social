package server

import (
	"fmt"
	"github.com/rs/cors"
	muxtrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gorilla/mux"
	"log/slog"
	"net/http"
	bhandler "ps-cats-social/pkg/base/handler"
	"time"
)

type Server struct {
	baseHandler *bhandler.BaseHTTPHandler
	router      *muxtrace.Router
	port        int
}

func NewServer(
	bHandler *bhandler.BaseHTTPHandler,
) Server {
	return Server{
		baseHandler: bHandler,
		port:        9093,
	}
}

func (s *Server) Run() error {
	slog.Info(fmt.Sprintf("Starting HTTP server at :%d ...", s.port))
	s.setupRouter()
	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	}).Handler(s.router)

	srv := &http.Server{
		Handler:      handler,
		Addr:         fmt.Sprintf(":%d", s.port),
		WriteTimeout: 60 * time.Second,
		ReadTimeout:  60 * time.Second,
	}
	return srv.ListenAndServe()
}
