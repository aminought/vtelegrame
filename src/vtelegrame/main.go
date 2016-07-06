package main

import (
	"vtelegrame/api/telegram"
	"vtelegrame/api/vk"
	"vtelegrame/core"

	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("main")

func main() {
	config := core.Config{}
	config.Load("config.json")
	log.Info("Config loaded")
	messages := core.Messages{}
	messages.Load("messages.json")
	log.Info("Messages loaded")

	vkLinkFactory := vk.LinkFactory{}
	vkAPI := vk.API{LinkFactory: vkLinkFactory}

	tgLinkFactory := telegram.LinkFactory{}
	telegramAPI := telegram.API{LinkFactory: tgLinkFactory}

	application := core.Application{VKApi: &vkAPI, TelegramAPI: &telegramAPI,
		Config: &config, Messages: &messages}
	application.Run()
}
