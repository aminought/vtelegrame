package vk

import "testing"
import "github.com/stretchr/testify/assert"

func TestLoad(t *testing.T) {
	var user User
	var token = "token"
	var apiUser = APIUser{ID: 1, Name: "Name", LastName: "LastName"}
	user.Load(token, apiUser)
	assert.Equal(t, token, user.AccessToken())
	assert.Equal(t, apiUser.ID, user.ID())
	assert.Equal(t, apiUser.Name, user.Name())
	assert.Equal(t, apiUser.LastName, user.LastName())
}
