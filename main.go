package main

import (
	"log"

	"web-project/config"

	"web-project/app"
)

func main() {

	app := app.NewApp()
	log.Fatalln(app.Start(":" + config.Configs.App.Port))
}
