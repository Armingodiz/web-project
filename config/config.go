package config

import (
	"fmt"
	"log"

	"github.com/kelseyhightower/envconfig"
)

type (
	Config struct {
		App      App
		Database Database
	}

	App struct {
		Port string `envconfig:"PORT" default:"3000"`
	}

	Database struct {
		Host     string `envconfig:"DATABASE_HOST"`
		Port     int    `envconfig:"DATABASE_PORT"`
		User     string `envconfig:"DATABASE_USER"`
		Password string `envconfig:"DATABASE_PASSWORD"`
		DbName   string `envconfig:"DATABASE_DBNAME"`
		Extras   string `envconfig:"DATABASE_EXTRAS"`
		Driver   string `envconfig:"DATABASE_DRIVER" default:"postgres"`
	}
)

var Configs Config

func init() {
	err := envconfig.Process("", &Configs)
	fmt.Println(Configs)
	if err != nil {
		log.Fatal(err.Error())
	}
}
