package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Import serve.go to ensure serveCmd is defined
var _ = serveCmd

var rootCmd = &cobra.Command{
	Use:   "api_rest",
	Short: "Unified Transport Operations League API CLI",
	Long:  `Unified Transport Operations League API - Command Line Interface`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
