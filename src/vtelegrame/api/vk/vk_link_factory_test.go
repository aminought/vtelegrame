package vk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildAuthorizeLink(t *testing.T) {
	clientID := "1"
	factory := LinkFactory{}

	link, data := factory.BuildAuthorizeLink(clientID)
	expectedLink := "https://oauth.vk.com/authorize"

	assert.Equal(t, expectedLink, link)
	assert.Equal(t, 6, len(data))
	assert.Equal(t, clientID, data["client_id"])
	assert.Equal(t, "friends,status,wall,groups,messages,notifications,offline", data["scope"])
	assert.Equal(t, "http://oauth.vk.com/blank.html", data["redirect_uri"])
	assert.Equal(t, "page", data["display"])
	assert.Equal(t, "token", data["response_type"])
	assert.Equal(t, "5.52", data["v"])
}

func TestBuildGetUsersLink(t *testing.T) {
	factory := LinkFactory{}
	fields := []string{"uid", "first_name"}
	ids := []int{1, 2}
	nameCase := "nom"
	token := "token"

	link, data := factory.BuildGetUsersLink(ids, fields, nameCase, token)
	expectedLink := "https://api.vk.com/method/users.get"

	assert.Equal(t, expectedLink, link)
	assert.Equal(t, 5, len(data))
	assert.Equal(t, "1,2", data["ids"])
	assert.Equal(t, "uid,first_name", data["fields"])
	assert.Equal(t, nameCase, data["name_case"])
	assert.Equal(t, token, data["access_token"])
	assert.Equal(t, "5.52", data["v"])
}

func TestBuildGetDialogsLink(t *testing.T) {
	factory := LinkFactory{}
	offset := 0
	count := 200
	previewLength := 0
	unread := true
	token := "token"

	expectedLink := "https://api.vk.com/method/messages.getDialogs"
	link, data := factory.BuildGetDialogsLink(offset, count, previewLength, unread, token)

	assert.Equal(t, expectedLink, link)
	assert.Equal(t, 6, len(data))
	assert.Equal(t, "0", data["offset"])
	assert.Equal(t, "200", data["count"])
	assert.Equal(t, "0", data["preview_length"])
	assert.Equal(t, "1", data["unread"])
	assert.Equal(t, token, data["access_token"])
	assert.Equal(t, "5.52", data["v"])
}

func TestBuildSendMessageLink(t *testing.T) {
	factory := LinkFactory{}
	userID := 1
	message := "message"
	token := "token"

	expectedLink := "https://api.vk.com/method/messages.send"
	link, data := factory.BuildSendMessageLink(userID, message, token)

	assert.Equal(t, expectedLink, link)
	assert.Equal(t, 5, len(data))
	assert.Equal(t, "1", data["user_id"])
	assert.Equal(t, message, data["message"])
	assert.Equal(t, token, data["access_token"])
	assert.Equal(t, "5.52", data["v"])
}

func TestBuildGetMessagesLink(t *testing.T) {
	factory := LinkFactory{}
	offset := 0
	count := 200
	filters := 1
	previewLength := 0
	token := "token"

	expectedLink := "https://api.vk.com/method/messages.get"
	link, data := factory.BuildGetMessagesLink(offset, count, filters, previewLength, token)

	assert.Equal(t, expectedLink, link)
	assert.Equal(t, 6, len(data))
	assert.Equal(t, "0", data["offset"])
	assert.Equal(t, "200", data["count"])
	assert.Equal(t, "1", data["filters"])
	assert.Equal(t, "0", data["preview_length"])
	assert.Equal(t, token, data["access_token"])
	assert.Equal(t, "5.52", data["v"])
}

func TestBuildGetMessagesByIdLink(t *testing.T) {
	factory := LinkFactory{}
	ids := []int{1, 2}
	previewLength := 0
	token := "token"

	expectedLink := "https://api.vk.com/method/messages.getById"
	link, data := factory.BuildGetMessagesByIDLink(ids, previewLength, token)

	assert.Equal(t, expectedLink, link)
	assert.Equal(t, 4, len(data))
	assert.Equal(t, "1,2", data["message_ids"])
	assert.Equal(t, "0", data["preview_length"])
	assert.Equal(t, token, data["access_token"])
	assert.Equal(t, "5.52", data["v"])
}

func TestBuildMarkMessageAsReadLink(t *testing.T) {
	factory := LinkFactory{}
	ids := []int{1, 2}
	token := "token"

	expectedLink := "https://api.vk.com/method/messages.markAsRead"
	link, data := factory.BuildMarkMessageAsReadLink(ids, token)

	assert.Equal(t, expectedLink, link)
	assert.Equal(t, 3, len(data))
	assert.Equal(t, "1,2", data["message_ids"])
	assert.Equal(t, token, data["access_token"])
	assert.Equal(t, "5.52", data["v"])
}
