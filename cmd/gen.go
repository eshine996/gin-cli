package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	genCmd.AddCommand(daoCmd)
	rootCmd.AddCommand(genCmd)
}

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "auto generate dao/service/controller",
}

var daoCmd = &cobra.Command{
	Use:   "dao",
	Short: "auto generate dao",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("done!")
	},
}
