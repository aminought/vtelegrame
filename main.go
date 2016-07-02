package main

import (
	"vtelegrame/api"
	"vtelegrame/app"

	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("main")

func main() {
	config := app.Config{}
	config.Load("config.json")
	log.Info("Config loaded")

	vkAPI := api.VKAPI{}
	telegramAPI := api.TelegramAPI{}
	application := app.Application{VKApi: &vkAPI, TelegramAPI: &telegramAPI, Config: &config}
	application.Run()
}
