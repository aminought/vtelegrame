package vk

import "testing"
import "github.com/stretchr/testify/assert"

func TestBuildAuthorizeLink(t *testing.T) {
	clientID := "1"
	factory := LinkFactory{}
	link := factory.BuildAuthorizeLink(clientID)
	expectedLink := "https://oauth.vk.com/authorize?client_id=1"
	expectedLink += "&scope=friends,status,wall,groups,messages,notifications,offline"
	expectedLink += "&redirect_uri=http://oauth.vk.com/blank.html"
	expectedLink += "&display=page&response_type=token"
	expectedLink += "&v=" + version
	assert.Equal(t, expectedLink, link)
}

func TestBuildGetUsersLink(t *testing.T) {
	factory := LinkFactory{}
	fields := []string{"uid", "first_name"}
	ids := []int{1, 2}
	nameCase := "nom"
	token := "token"
	link := factory.BuildGetUsersLink(ids, fields, nameCase, token)

	expectedLink := "https://api.vk.com/method/users.get?ids=1,2&fields=uid,first_name&"
	expectedLink += "name_case=nom&access_token=token&v=" + version

	assert.Equal(t, expectedLink, link)
}

func TestBuildGetDialogsLink(t *testing.T) {
	factory := LinkFactory{}
	offset := 0
	count := 200
	previewLength := 0
	unread := true
	token := "token"

	expectedLink := "https://api.vk.com/method/messages.getDialogs?offset=0"
	expectedLink += "&count=200&preview_length=0&unread=1&access_token=token"
	expectedLink += "&v=" + version
	link := factory.BuildGetDialogsLink(offset, count, previewLength, unread, token)

	assert.Equal(t, expectedLink, link)
}

func TestBuildSendMessageLink(t *testing.T) {
	factory := LinkFactory{}
	expectedLink := "https://api.vk.com/method/messages.send"
	link := factory.BuildSendMessageLink()

	assert.Equal(t, expectedLink, link)
}

func TestBuildGetMessagesLink(t *testing.T) {
	factory := LinkFactory{}
	offset := 0
	count := 200
	filters := 1
	previewLength := 0
	token := "token"

	expectedLink := "https://api.vk.com/method/messages.get?offset=0"
	expectedLink += "&count=200&filters=1&preview_length=0&access_token=token"
	expectedLink += "&v=" + version
	link := factory.BuildGetMessagesLink(offset, count, filters, previewLength, token)

	assert.Equal(t, expectedLink, link)
}

func TestBuildGetMessagesByIdLink(t *testing.T) {
	factory := LinkFactory{}
	ids := []int{1, 2}
	previewLength := 0
	token := "token"

	expectedLink := "https://api.vk.com/method/messages.getById?message_ids=1,2"
	expectedLink += "&preview_length=0&access_token=token&v=" + version
	link := factory.BuildGetMessagesByIDLink(ids, previewLength, token)

	assert.Equal(t, expectedLink, link)
}

func TestBuildMarkMessageAsReadLink(t *testing.T) {
	factory := LinkFactory{}
	ids := []int{1, 2}
	token := "token"

	expectedLink := "https://api.vk.com/method/messages.markAsRead?message_ids=1,2"
	expectedLink += "&access_token=token&v=" + version
	link := factory.BuildMarkMessageAsReadLink(ids, token)

	assert.Equal(t, expectedLink, link)
}
