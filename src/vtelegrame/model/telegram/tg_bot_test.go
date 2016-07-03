package telegram

import "testing"
import "github.com/stretchr/testify/assert"

func TestLoad(t *testing.T) {
	var bot Bot
	token := "token"
	id := 1
	username := "username"
	bot.Load(token, id, username)
	assert.Equal(t, token, bot.AccessToken())
	assert.Equal(t, id, bot.ID())
	assert.Equal(t, username, bot.Username())
}
