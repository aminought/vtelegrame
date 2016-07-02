package vk

import "testing"
import "github.com/stretchr/testify/assert"

func TestBuildAuthorizeLink(t *testing.T) {
	clientID := "1"
	factory := LinkFactory{}
	link := factory.BuildAuthorizeLink(clientID)
	expectedLink := "https://oauth.vk.com/authorize?client_id=1&"
	expectedLink += "scope=friends,status,wall,groups,messages,notifications,offline&"
	expectedLink += "redirect_uri=http://oauth.vk.com/blank.html&"
	expectedLink += "display=page&response_type=token&v=5.52"
	assert.Equal(t, expectedLink, link)
}

func TestBuildGetUsersLink(t *testing.T) {
	factory := LinkFactory{}
	fields := []string{"uid", "first_name"}
	nameCase := "nom"
	token := "token"
	link := factory.BuildGetUsersLink(nil, fields, nameCase, token)

	expectedLink := "https://api.vk.com/method/users.get?fields=uid,first_name&"
	expectedLink += "name_case=nom&access_token=token"

	assert.Equal(t, expectedLink, link)
}
