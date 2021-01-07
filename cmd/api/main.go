package main

import (
	"github.com/apex/log"

	"github.com/brauliodev29/go-guest/config"
	"github.com/brauliodev29/go-guest/server"
)

func main() {
	if err := config.Init(); err != nil {
		log.WithError(err)
	}

	app, err := server.NewApp()
	if err != nil {
		log.WithError(err)
	}

	if err := app.Run(config.Config.Port); err != nil {
		log.WithError(err)
	}
}
