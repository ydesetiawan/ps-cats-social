package server

import (
	"fmt"
	"log/slog"
	"net/http"
	catbandler "ps-cats-social/internal/cat/handler"
	"ps-cats-social/internal/shared"
	userhandler "ps-cats-social/internal/user/handler"
	bhandler "ps-cats-social/pkg/base/handler"
	"time"

	"github.com/rs/cors"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"
	muxtrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gorilla/mux"
)

type Server struct {
	baseHandler     *bhandler.BaseHTTPHandler
	userHandler     *userhandler.UserHTTPHandler
	catHandler      *catbandler.CatHttpHandler
	catMatchHandler *catbandler.CatMatchHTTPHandler
	router          *muxtrace.Router
	port            int
}

func NewServer(
	bHandler *bhandler.BaseHTTPHandler,
	userHandler *userhandler.UserHTTPHandler,
	catHandler *catbandler.CatHttpHandler,
	catMatchHandler *catbandler.CatMatchHTTPHandler,
	port int,
) Server {
	return Server{
		baseHandler:     bHandler,
		userHandler:     userHandler,
		catHandler:      catHandler,
		catMatchHandler: catMatchHandler,
		router:          muxtrace.NewRouter(muxtrace.WithServiceName(shared.ServiceName)),
		port:            port,
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
