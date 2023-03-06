package main

import (
	"os"
	"web/api"
	"web/app"
	"web/common"
)

func main() {
	data, _ := os.ReadFile("application.yml")
	common.ParseAppConfig(data)
	api.Init()
	app.Run()
}
