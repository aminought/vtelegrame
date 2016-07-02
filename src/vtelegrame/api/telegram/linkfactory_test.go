package telegram

import "testing"
import "github.com/stretchr/testify/assert"

func TestBuildGetMeLink(t *testing.T) {
	factory := LinkFactory{}
	token := "token"
	link := factory.BuildGetMeLink(token)
	expectedLink := "https://api.telegram.org/bottoken/getMe"
	assert.Equal(t, expectedLink, link)
}

func TestBuildSendMessageLink(t *testing.T) {
	factory := LinkFactory{}
	token := "token"
	link := factory.BuildSendMessageLink(token)
	expectedLink := "https://api.telegram.org/bottoken/sendMessage"
	assert.Equal(t, expectedLink, link)
}

func TestBuildGetUpdatesLink(t *testing.T) {
	factory := LinkFactory{}
	token := "token"
	link := factory.BuildGetUpdatesLink(token)
	expectedLink := "https://api.telegram.org/bottoken/getUpdates"
	assert.Equal(t, expectedLink, link)
}
