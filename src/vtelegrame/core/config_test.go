package core

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	var config Config
	file, err := ioutil.TempFile(os.TempDir(), "config")
	if err != nil {
		t.Fatal("Cannot to create temporary file: " + err.Error())
	}
	defer os.Remove(file.Name())

	file.WriteString(`{"client_id": "1", "access_token_vk": "vk", 
        "access_token_telegram": "telegram", "telegram_user": "user"}`)

	config.Load(file.Name())
	assert.Equal(t, config.ClientID, "1")
	assert.Equal(t, config.VKAccessToken, "vk")
	assert.Equal(t, config.TelegramAccessToken, "telegram")
	assert.Equal(t, config.TelegramUser, "user")
}
