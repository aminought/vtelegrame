package vk

// User is a class representing user profile in vk.com
type User struct {
	id          int
	name        string
	lastName    string
	accessToken string
}

// ID returns vk user id
func (user *User) ID() int {
	return user.id
}

// Name returns vk user name
func (user *User) Name() string {
	return user.name
}

// LastName returns vk user last name
func (user *User) LastName() string {
	return user.lastName
}

// AccessToken returns vk user access token
func (user *User) AccessToken() string {
	return user.accessToken
}

// Load user info
func (user *User) Load(token string, apiUser APIUser) {
	user.accessToken = token
	user.id = apiUser.ID
	user.name = apiUser.Name
	user.lastName = apiUser.LastName
}
