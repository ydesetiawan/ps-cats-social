package main

import (
	"fmt"
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

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"golang.org/x/exp/slog"
)

var port int

var httpCmd = &cobra.Command{
	Use:   "http [OPTIONS]",
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

func init() {
	httpCmd.Flags().IntVarP(&port, "port", "p", 8080, "Port to run the HTTP server")
}

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
	{

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

	httpServer := server.NewServer(
		baseHandler, userHandler, catHandler, catMatchHandler, port,
	)
	return httpServer.Run()
}

func dbInitConnection() *sqlx.DB {
	godotenv.Load(".env")

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	uname := os.Getenv("DB_USERNAME")
	pass := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	dbparams := os.Getenv("DB_PARAMS")

	return psqlqgen.Init(host, port, uname, pass, dbname, dbparams, shared.ServiceName)
}

func initInfra() {
	dbConnection := dbInitConnection()

	userRepository := userrepository.NewUserRepositoryImpl(dbConnection)
	catRepository := catrepository.NewCatRepositoryImpl(dbConnection)
	catMatchRepository := catrepository.NewCatMatchRepositoryImpl(dbConnection)

	userService := userservice.NewUserService(userRepository)
	catService := catservice.NewCatService(catRepository, catMatchRepository)
	catMatchService := catservice.NewCatMatchService(catRepository, userRepository, catMatchRepository)

	catHandler = cathandler.NewCatHttpHandler(catService)
	userHandler = userhandler.NewUserHTTPHandler(userService)
	catMatchHandler = cathandler.NewCatMatchHTTPHandler(catMatchService)
}
