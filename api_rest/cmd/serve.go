package cmd

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/UTOL-s/module/api_rest/internal"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// setupSignalHandling sets up graceful shutdown on system signals
func setupSignalHandling() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigChan
		zap.L().Info("received shutdown signal", zap.String("signal", sig.String()))
		os.Exit(0)
	}()
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the API server",
	Long:  `Start the Unified Transport Operations League API server`,
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

func serve() {
	// Create the FX application
	app := internal.Bootstrap()

	// Start the application
	ctx := context.Background()
	if err := app.Start(ctx); err != nil {
		log.Fatal("Failed to start application:", err)
	}

	// Wait for interrupt signal
	<-app.Done()

	// Stop the application gracefully
	stopCtx, cancel := context.WithTimeout(ctx, 30)
	defer cancel()

	if err := app.Stop(stopCtx); err != nil {
		log.Fatal("Failed to stop application:", err)
	}
}
