package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gin-cli",
	Short: "Engineering example based on gin framework",
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
