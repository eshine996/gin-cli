package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of gin-cli",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gin-cli version: 1.0.0")
	},
}
