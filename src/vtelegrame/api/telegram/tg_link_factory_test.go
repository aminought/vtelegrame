package telegram

import "testing"
import "github.com/stretchr/testify/assert"

var start = "https://api.telegram.org/bot"

func TestBuildGetMeLink(t *testing.T) {
	factory := LinkFactory{}
	token := "token"
	link := factory.BuildGetMeLink(token)
	expectedLink := start + "token/getMe"
	assert.Equal(t, expectedLink, link)
}

func TestBuildSendMessageLink(t *testing.T) {
	factory := LinkFactory{}
	chatID := 1
	text := "text"
	token := "token"
	link, data := factory.BuildSendMessageLink(chatID, text, token)

	expectedLink := start + "token/sendMessage"

	assert.Equal(t, expectedLink, link)
	assert.Equal(t, 2, len(data))
	assert.Equal(t, "1", data["chat_id"])
	assert.Equal(t, text, data["text"])
}

func TestBuildGetUpdatesLink(t *testing.T) {
	factory := LinkFactory{}
	token := "token"
	link := factory.BuildGetUpdatesLink(token)
	expectedLink := "https://api.telegram.org/bottoken/getUpdates"
	assert.Equal(t, expectedLink, link)
}
