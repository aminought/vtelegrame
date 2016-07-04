package telegram

import (
	"encoding/json"
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
	answer := http.GetRequest(link, nil)

	result := getMeResult{}
	json.Unmarshal(answer, &result)
	var bot = new(telegram.Bot)
	bot.Load(token, result.User.ID, result.User.UserName)
	return bot
}

// SendMessage sends text message to concrete user
func (api *API) SendMessage(bot *telegram.Bot, text string) {
	link, data := api.LinkFactory.BuildSendMessageLink(bot.ChatID(), text, bot.AccessToken())
	http.PostRequest(link, data)
}

type getUpdatesResult struct {
	Updates []telegram.Update `json:"result"`
}

// GetUpdates return bot updates
func (api *API) GetUpdates(bot *telegram.Bot) []telegram.Update {
	link := api.LinkFactory.BuildGetUpdatesLink(bot.AccessToken())
	answer := http.GetRequest(link, nil)

	result := getUpdatesResult{}
	json.Unmarshal(answer, &result)
	return result.Updates
}
