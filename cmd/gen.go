package cmd

import (
	"fmt"
	"github.com/codeHauler-1/gin-cli/app/gen"
	g "github.com/codeHauler-1/gin-cli/global"
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
		dns := g.Config.GetString("Mysql.default.link")
		if dns == "" {
			fmt.Println("Error:not found mysql link in config.yaml")
			return
		}

		if err := gen.GenerateDao(dns); err != nil {
			fmt.Println(fmt.Sprintf("Error:%s", err.Error()))
		}
		fmt.Println("done!")
	},
}
