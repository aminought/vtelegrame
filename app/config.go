package app

import (
	"io/ioutil"

	"encoding/json"

	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("app")

// Config struct
type Config struct {
	ClientID    string `json:"client_id"`
	AccessToken string `json:"access_token"`
}

// Load configuration
func (config *Config) Load(path string) {
	buf, _ := ioutil.ReadFile(path)
	json.Unmarshal(buf, config)
}
