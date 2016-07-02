package main

import (
	"vtelegrame/api/telegram"
	"vtelegrame/api/vk"
	"vtelegrame/app"

	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("main")

func main() {
	config := app.Config{}
	config.Load("config.json")
	log.Info("Config loaded")

	vkLinkFactory := vk.LinkFactory{}
	vkAPI := vk.API{LinkFactory: vkLinkFactory}

	tgLinkFactory := telegram.LinkFactory{}
	telegramAPI := telegram.API{LinkFactory: tgLinkFactory}

	application := app.Application{VKApi: &vkAPI, TelegramAPI: &telegramAPI, Config: &config}
	application.Run()
}
