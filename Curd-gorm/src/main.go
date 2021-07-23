package main

import (
	"github.com/emisgroup/request-annotator/api/app"
	"github.com/emisgroup/request-annotator/api/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3003")
}