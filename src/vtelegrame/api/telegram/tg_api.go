package telegram

import (
	"encoding/json"
	"strconv"
	"vtelegrame/http"
	"vtelegrame/model/telegram"
)

// API represents api for telegram.org
type API struct {
	LinkFactory LinkFactory
}

type getMeResult struct {
	User telegram.User `json:"result"`
}

// GetMe returns Bot object
func (api *API) GetMe(token string) *telegram.Bot {
	link := api.LinkFactory.BuildGetMeLink(token)
	data := http.GetRequest(link)
	result := getMeResult{}
	json.Unmarshal(data, &result)
	var bot = new(telegram.Bot)
	bot.Load(token, result.User.ID, result.User.UserName)
	return bot
}

// SendMessage sends text message to concrete user
func (api *API) SendMessage(bot *telegram.Bot, text string) {
	link := api.LinkFactory.BuildSendMessageLink(bot.AccessToken())
	data := make(map[string]string)
	data["chat_id"] = strconv.Itoa(bot.ChatID())
	data["text"] = text
	http.PostRequest(link, data)
}

type getUpdatesResult struct {
	Updates []telegram.Update `json:"result"`
}

// GetUpdates return bot updates
func (api *API) GetUpdates(bot *telegram.Bot) []telegram.Update {
	link := api.LinkFactory.BuildGetUpdatesLink(bot.AccessToken())
	data := http.GetRequest(link)
	result := getUpdatesResult{}
	json.Unmarshal(data, &result)
	return result.Updates
}
