package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gin-cli",
	Short: "gin framework client helper",
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
