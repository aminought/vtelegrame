package core

import (
	"io/ioutil"

	"encoding/json"

	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("app")

// Config struct
type Config struct {
	ClientID            string `json:"client_id"`
	VKAccessToken       string `json:"access_token_vk"`
	TelegramAccessToken string `json:"access_token_telegram"`
	TelegramUser        string `json:"telegram_user"`
}

// Load configuration
func (config *Config) Load(path string) {
	buf, _ := ioutil.ReadFile(path)
	json.Unmarshal(buf, config)
}
