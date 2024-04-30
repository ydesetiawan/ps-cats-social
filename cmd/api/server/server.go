package server

import (
	"fmt"
	"github.com/rs/cors"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"
	muxtrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gorilla/mux"
	"log/slog"
	"net/http"
	"ps-cats-social/internal/shared"
	"ps-cats-social/internal/user/handler"
	bhandler "ps-cats-social/pkg/base/handler"
	"time"
)

type Server struct {
	baseHandler *bhandler.BaseHTTPHandler
	userHandler *handler.UserHTTPHandler
	router      *muxtrace.Router
	port        int
}

func NewServer(
	bHandler *bhandler.BaseHTTPHandler,
	userHandler *handler.UserHTTPHandler,
) Server {
	return Server{
		baseHandler: bHandler,
		userHandler: userHandler,
		router:      muxtrace.NewRouter(muxtrace.WithServiceName(shared.ServiceName)),
		port:        8080,
	}
}

func (s *Server) Run() error {
	slog.Info(fmt.Sprintf("Starting HTTP server at :%d ...", s.port))
	s.router.Use(otelmux.Middleware(shared.ServiceName))
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
