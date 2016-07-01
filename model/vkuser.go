package model

// VKUser is a class representing user profile in vk.com
type VKUser struct {
	id          int
	name        string
	lastName    string
	accessToken string
}

// ID returns vk user id
func (vkuser *VKUser) ID() int {
	return vkuser.id
}

// Name returns vk user name
func (vkuser *VKUser) Name() string {
	return vkuser.name
}

// LastName returns vk user last name
func (vkuser *VKUser) LastName() string {
	return vkuser.lastName
}

// AccessToken returns vk user access token
func (vkuser *VKUser) AccessToken() string {
	return vkuser.accessToken
}

// Load user info
func (vkuser *VKUser) Load(token string, vkAPIUser VKAPIUser) {
	vkuser.accessToken = token
	vkuser.id = vkAPIUser.ID
	vkuser.name = vkAPIUser.Name
	vkuser.lastName = vkAPIUser.LastName
}
