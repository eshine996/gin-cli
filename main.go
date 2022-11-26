package main

import "github.com/codeHauler-1/gin-cli/cmd"

func main() {
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
