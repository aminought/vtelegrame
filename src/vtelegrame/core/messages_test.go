package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetVkAutoAnswer(t *testing.T) {
	messages := Messages{}
	messages.VkAutoAnswer = []string{"Ты меня еще не добавил%sex%: telegram.me/%tg_login%?"}

	message := messages.GetVkAutoAnswer(female, "login")
	expected := "Ты меня еще не добавила: telegram.me/login?"

	assert.Equal(t, expected, message)
}

func TestGetTgMessageFormat(t *testing.T) {
	messages := Messages{}
	messages.TgMessageFormat = []string{"%firstname% %lastname% %chat_braces%:", "", "%text%"}

	message := messages.GetTgMessageFormat("name", "surname", "title", "message")
	expected := "name surname (title):\n\nmessage"

	assert.Equal(t, expected, message)
}

func TestGetTgStartMessageFormat(t *testing.T) {
	messages := Messages{}
	messages.TgStartMessage = "Hello, %tg_login%."
	message := messages.GetTgStartMessageFormat("login")
	expected := "Hello, login."
	assert.Equal(t, expected, message)
}
