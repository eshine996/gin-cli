package main

import "gin-cli/cmd"

func main() {
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
