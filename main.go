package main

import (
	"vtelegrame/api"
	"vtelegrame/app"

	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("main")
var config = app.Config{}

func main() {
	config.Load("config.json")
	log.Info(config.ClientID)
	vkAPI := api.VKAPI{ClientID: config.ClientID}
	application := app.Application{VKApi: &vkAPI}
	application.Run()
}
