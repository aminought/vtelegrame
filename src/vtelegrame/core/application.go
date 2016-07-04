package core

import (
	"errors"
	"fmt"
	tgAPI "vtelegrame/api/telegram"
	vkAPI "vtelegrame/api/vk"
	tgModel "vtelegrame/model/telegram"
	vkModel "vtelegrame/model/vk"

	"github.com/jasonlvhit/gocron"
)

// Application is a main struct
type Application struct {
	VKApi       *vkAPI.API
	TelegramAPI *tgAPI.API
	Config      *Config
	VKUser      *vkModel.User
	Bot         *tgModel.Bot
}

// Run is an entry point of application
func (application *Application) Run() {
	tokenVK := application.getVKAccessToken()
	application.loadVKUser(tokenVK)
	tokenTelegram := application.getTelegramAccessToken()
	application.loadTelegramBot(tokenTelegram)
	application.sendHelloMessage()
	application.startHandleVKMessageTask()
}

func (application *Application) getVKAccessToken() string {
	if application.Config.VKAccessToken != "" {
		return application.Config.VKAccessToken
	}
	log.Panic("vk.com access token is empty")
	return ""
}

func (application *Application) getTelegramAccessToken() string {
	return application.Config.TelegramAccessToken
}

func (application *Application) loadVKUser(token string) {
	var fields = []string{"uid", "first_name", "last_name"}
	users := application.VKApi.GetUsers(nil, fields, "nom", token)
	if len(users) > 0 {
		application.VKUser = new(vkModel.User)
		user := application.VKUser
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

func (application *Application) sendHelloMessage() {
	config := application.Config
	target := config.TelegramUser
	var text = "Hello, " + target + "! Bot will be ready for crossposting soon."
	updates := application.TelegramAPI.GetUpdates(application.Bot)
	chatID, err := application.getChatID(updates, target)
	if err != nil {
		log.Panic(err.Error())
	}

	application.Bot.SetChatID(chatID)
	application.TelegramAPI.SendMessage(application.Bot, text)
	log.Info("Message \"" + text + "\" was sent")
}

func (application *Application) getChatID(updates []tgModel.Update, username string) (int, error) {
	for _, update := range updates {
		if update.Message.From.UserName == username {
			return update.Message.Chat.ID, nil
		}
	}
	return -1, errors.New("Invalid user name")
}

func (application *Application) startHandleVKMessageTask() {
	log.Info("Configure cron")
	gocron.Every(2).Seconds().Do(HandleVKMessages, application.VKApi,
		application.TelegramAPI, application.VKUser.AccessToken(), application.Bot)
	log.Info("Starting cron")
	<-gocron.Start()
}
