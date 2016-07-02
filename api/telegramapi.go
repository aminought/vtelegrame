package api

import (
	"vtelegrame/http"
	"vtelegrame/model/telegram"
)

import "encoding/json"

const telegramAPIRequestStart = "https://api.telegram.org/bot"

// TelegramAPI represents api for telegram.org
type TelegramAPI struct{}

type getMeResult struct {
	User telegram.User `json:"result"`
}

// GetMe returns Bot object
func (telegramAPI *TelegramAPI) GetMe(token string) telegram.Bot {
	link := telegramAPIRequestStart + token + "/getMe"
	data := http.GetRequest(link)
	result := getMeResult{}
	json.Unmarshal(data, &result)
	var bot = telegram.Bot{}
	bot.Load(token, result.User.ID, result.User.UserName)
	return bot
}

// SendMessage sends text message to concrete user
func (telegramAPI *TelegramAPI) SendMessage(bot telegram.Bot, target string, text string) {
	link := telegramAPIRequestStart + bot.AccessToken() + "/sendMessage?"
	var data = make(map[string]string)
	data["chat_id"] = target
	data["text"] = text
	http.PostRequest(link, data)
}

type getUpdatesResult struct {
	Updates []telegram.Update `json:"result"`
}

// GetUpdates return bot updates
func (telegramAPI *TelegramAPI) GetUpdates(bot telegram.Bot) []telegram.Update {
	link := telegramAPIRequestStart + bot.AccessToken() + "/getUpdates"
	data := http.GetRequest(link)
	result := getUpdatesResult{}
	json.Unmarshal(data, &result)
	return result.Updates
}
