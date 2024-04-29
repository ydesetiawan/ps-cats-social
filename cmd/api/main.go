package api

import (
	"fmt"
	"github.com/spf13/cobra"
	"log/slog"
	"os"
	"ps-cats-social/cmd/api/server"
	bhandler "ps-cats-social/pkg/base/handler"
)

var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "Run HTTP API",
	Long:  "Run HTTP API for SCM",
	RunE:  runHttpCommand,
}

var (
	params      map[string]string
	baseHandler *bhandler.BaseHTTPHandler
)

func main() {
	if err := httpCmd.Execute(); err != nil {
		slog.Error(fmt.Sprintf("Error on command execution: %s", err.Error()))
		os.Exit(1)
	}
}

func runHttpCommand(cmd *cobra.Command, args []string) error {

	httpServer := server.NewServer(
		baseHandler,
	)
	return httpServer.Run()
}
