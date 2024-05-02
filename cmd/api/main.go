package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/cobra"
	"go.opentelemetry.io/otel"
	"golang.org/x/exp/slog"
	ddotel "gopkg.in/DataDog/dd-trace-go.v1/ddtrace/opentelemetry"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"gopkg.in/DataDog/dd-trace-go.v1/profiler"
	stdlog "log"
	"os"
	"ps-cats-social/cmd/api/server"
	cathandler "ps-cats-social/internal/cat/handler"
	catrepository "ps-cats-social/internal/cat/repository"
	catservice "ps-cats-social/internal/cat/service"
	"ps-cats-social/internal/shared"
	userhandler "ps-cats-social/internal/user/handler"
	userrepository "ps-cats-social/internal/user/repository"
	userservice "ps-cats-social/internal/user/service"
	bhandler "ps-cats-social/pkg/base/handler"
	"ps-cats-social/pkg/logger"
	psqlqgen "ps-cats-social/pkg/psqlqgen"
	"strings"
	"time"
)

var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "Run HTTP API",
	Long:  "Run HTTP API for SCM",
	RunE:  runHttpCommand,
}

var (
	params          map[string]string
	baseHandler     *bhandler.BaseHTTPHandler
	userHandler     *userhandler.UserHTTPHandler
	catHandler      *cathandler.CatHttpHandler
	catMatchHandler *cathandler.CatMatchHTTPHandler
)

func main() {
	if err := httpCmd.Execute(); err != nil {
		slog.Error(fmt.Sprintf("Error on command execution: %s", err.Error()))
		os.Exit(1)
	}
}

func logLevel() slog.Level {
	switch strings.ToLower(os.Getenv("LOG_LEVEL")) {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

func initLogger() {
	// STEP 1: Prepare Logger to use our standard logger.
	{
		// Prepare logger and set it as global logger within the application.
		// Use block statement to ensure we don't use `logger` variable directly.
		log, err := logger.SlogOption{
			Resource: map[string]string{
				"service.name":        shared.ServiceName,
				"service.ns":          "cats_social",
				"service.instance_id": "random-uuid",
				"service.version":     "v.0",
				"service.env":         "staging",
			},
			ContextExtractor:   nil,
			AttributeFormatter: nil,
			Writer:             os.Stdout,
			Leveler:            logLevel(),
		}.NewSlog()
		if err != nil {
			err = fmt.Errorf("prepare logger error: %w", err)
			stdlog.Fatal(err) // if logger cannot be prepared (commonly due to option value error), use std logger.
			return
		}

		// Set logger as global logger.
		slog.SetDefault(log)
	}
}

func runHttpCommand(cmd *cobra.Command, args []string) error {
	initLogger()
	initInfra()
	//server.InitDBMigrate()
	// init datadog tracer
	rules := []tracer.SamplingRule{tracer.RateRule(1)}
	tracerOpt := []tracer.StartOption{
		tracer.WithRuntimeMetrics(),
		tracer.WithTraceEnabled(true),
		tracer.WithService(shared.ServiceName),
		tracer.WithEnv("shared.DatadogEnv"),
		tracer.WithSamplingRules(rules),
	}
	tracerProvider := ddotel.NewTracerProvider(tracerOpt...)
	defer func() {
		if tracerProvider == nil {
			return
		}
		if _err := tracerProvider.Shutdown(); _err != nil {
			slog.Error("OpenTelemetry provider shutdown error", slog.Any("error", _err))
		}
	}()
	otel.SetTracerProvider(tracerProvider)
	// init datadog profiler
	if os.Getenv("DD_USE_PROFILER") == "true" {
		if err := profiler.Start(
			profiler.WithService(shared.ServiceName),
			profiler.WithEnv(shared.DatadogEnv),
			profiler.WithPeriod(time.Second),
		); err != nil {
			slog.Warn(fmt.Sprintf("error start profiler: %s", err.Error()))
		}
		defer profiler.Stop()
	}

	httpServer := server.NewServer(
		baseHandler, userHandler, catHandler, catMatchHandler,
	)
	return httpServer.Run()
}

func dbInitConnection() *sqlx.DB {
	//host := os.Getenv("DB_HOST")
	//port := os.Getenv("DB_PORT")
	//uname := os.Getenv("DB_USER")
	//pass := os.Getenv("DB_PASSWORD")
	//dbname := os.Getenv("DB_NAME")

	host := "localhost"
	port := "5432"
	uname := "postgres"
	pass := "123"
	dbname := "cats_social"

	return psqlqgen.Init(host, port, uname, pass, dbname, shared.ServiceName)
}

func initInfra() {
	dbConnection := dbInitConnection()

	userRepository := userrepository.NewUserRepositoryImpl(dbConnection)
	userService := userservice.NewUserService(userRepository)
	userHandler = userhandler.NewUserHTTPHandler(userService)

	catRepository := catrepository.NewCatRepositoryImpl(dbConnection)
	catService := catservice.NewCatService(catRepository)
	catHandler = cathandler.NewCatHttpHandler(catService)

	catMatchRepository := catrepository.NewCatMatchRepositoryImpl(dbConnection)
	catMatchService := catservice.NewCatMatchService(catRepository, userRepository, catMatchRepository)
	catMatchHandler = cathandler.NewCatMatchHTTPHandler(catMatchService)
}
