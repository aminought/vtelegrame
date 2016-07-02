package api

import "vtelegrame/model"
import "vtelegrame/http"
import "encoding/json"

const telegramAPIRequestStart = "https://api.telegram.org/bot"

// TelegramAPI represents api for telegram.org
type TelegramAPI struct{}

type userJSONContainer struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserName  string `json:"username"`
}

type getMeResultJSONContainer struct {
	User userJSONContainer `json:"result"`
}

// GetMe returns Bot object
func (telegramAPI *TelegramAPI) GetMe(token string) model.Bot {
	link := telegramAPIRequestStart + token + "/getMe"
	data := http.Request(link)
	result := getMeResultJSONContainer{}
	json.Unmarshal(data, &result)
	var bot = model.Bot{}
	bot.Load(token, result.User.ID, result.User.UserName)
	return bot
}
