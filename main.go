package main

import "git.xiaoyanggroup.cn/xyyjyframework/gin-cli/cmd"

func main() {
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
