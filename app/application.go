package app

import (
	"fmt"
	"vtelegrame/api"
	"vtelegrame/model"

	"github.com/skratchdot/open-golang/open"
)

// Application is a main struct
type Application struct {
	VKApi       *api.VKAPI
	TelegramAPI *api.TelegramAPI
	Config      *Config
	VkUser      model.VKUser
	Bot         model.Bot
}

// Run is an entry point of application
func (application *Application) Run() {
	tokenVK := application.getVKAccessToken()
	application.loadVKUser(tokenVK)
	tokenTelegram := application.getTelegramAccessToken()
	application.loadTelegramBot(tokenTelegram)
}

func (application *Application) getVKAccessToken() string {
	if application.Config.VKAccessToken != "" {
		return application.Config.VKAccessToken
	}

	fmt.Println("Hello, user. Let's authorize in vk.com.")
	var link = application.VKApi.GetAuthorizeLink()
	open.Start(link)

	fmt.Println("Please, insert here your access token:")
	var token string
	fmt.Scanf("%s", &token)

	return token
}

func (application *Application) getTelegramAccessToken() string {
	return application.Config.TelegramAccessToken
}

func (application *Application) loadVKUser(token string) {
	var fields = []string{"uid", "first_name", "last_name"}
	users := application.VKApi.GetUsers(nil, fields, "nom", token)
	if len(users) > 0 {
		var user = application.VkUser
		user.Load(token, users[0])
		fmt.Println("Hello, " + user.Name() + " " + user.LastName() + "!")
	} else {
		log.Error("Who are you?")
	}
}

func (application *Application) loadTelegramBot(token string) {
	application.Bot = application.TelegramAPI.GetMe(token)
	log.Info("Bot loaded")
}
