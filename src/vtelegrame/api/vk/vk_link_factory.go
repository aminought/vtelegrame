package vk

import (
	"math/rand"
	"strconv"
	"vtelegrame/util"
)

// LinkFactory is a link factory for vk.com API
type LinkFactory struct{}

const apiRequestStart = "https://api.vk.com/method/"
const version = "5.52"

// BuildAuthorizeLink returns link for authorization
func (factory *LinkFactory) BuildAuthorizeLink(clientID string) (string, map[string]string) {
	link := "https://oauth.vk.com/authorize"
	data := make(map[string]string)
	data["client_id"] = clientID
	data["scope"] = "friends,status,wall,groups,messages,notifications,offline"
	data["redirect_uri"] = "http://oauth.vk.com/blank.html"
	data["display"] = "page"
	data["response_type"] = "token"
	data["v"] = version
	return link, data
}

// BuildGetUsersLink returns link for users.get method
func (factory *LinkFactory) BuildGetUsersLink(ids []int, fields []string,
	nameCase string, token string) (string, map[string]string) {
	const method = "users.get"

	idArr := util.IntArray(ids)
	fieldArr := util.StringArray(fields)

	link := apiRequestStart + method
	data := make(map[string]string)
	data["ids"] = util.GetCommaSeparatedString(&idArr)
	data["fields"] = util.GetCommaSeparatedString(&fieldArr)
	data["name_case"] = nameCase
	data["access_token"] = token
	data["v"] = version

	return link, data
}

// BuildGetDialogsLink returns link for messages.getDialogs method
func (factory *LinkFactory) BuildGetDialogsLink(offset int, count int,
	previewLength int, unread bool, token string) (string, map[string]string) {

	const method = "messages.getDialogs"
	link := apiRequestStart + method
	data := make(map[string]string)
	data["offset"] = strconv.Itoa(offset)
	data["count"] = strconv.Itoa(count)
	data["preview_length"] = strconv.Itoa(previewLength)
	u := "1"
	if !unread {
		u = "0"
	}
	data["unread"] = u
	data["access_token"] = token
	data["v"] = version

	return link, data
}

// BuildSendMessageLink returns link for mesages.send method
func (factory *LinkFactory) BuildSendMessageLink(userID int,
	message string, token string) (string, map[string]string) {
	const method = "messages.send"
	link := apiRequestStart + method
	data := make(map[string]string)
	data["user_id"] = strconv.Itoa(userID)
	data["random_id"] = strconv.Itoa(rand.Int())
	data["message"] = message
	data["access_token"] = token
	data["v"] = version
	return link, data
}

//BuildGetMessagesLink returns link for messages.get method
func (factory *LinkFactory) BuildGetMessagesLink(offset int,
	count int, filters int, previewLength int, token string) (string, map[string]string) {

	const method = "messages.get"
	link := apiRequestStart + method
	data := make(map[string]string)
	data["offset"] = strconv.Itoa(offset)
	data["count"] = strconv.Itoa(count)
	data["filters"] = strconv.Itoa(filters)
	data["preview_length"] = strconv.Itoa(previewLength)
	data["access_token"] = token
	data["v"] = version

	return link, data
}

// BuildGetMessagesByIDLink returns link for messages.getById method
func (factory *LinkFactory) BuildGetMessagesByIDLink(ids []int,
	previewLength int, token string) (string, map[string]string) {

	const method = "messages.getById"
	link := apiRequestStart + method
	idArr := util.IntArray(ids)
	data := make(map[string]string)
	data["message_ids"] = util.GetCommaSeparatedString(&idArr)
	data["preview_length"] = strconv.Itoa(previewLength)
	data["access_token"] = token
	data["v"] = version

	return link, data
}

// BuildMarkMessageAsReadLink returns link for messages.markAsRead method
func (factory *LinkFactory) BuildMarkMessageAsReadLink(ids []int,
	token string) (string, map[string]string) {
	const method = "messages.markAsRead"
	link := apiRequestStart + method
	idArr := util.IntArray(ids)
	data := make(map[string]string)
	data["message_ids"] = util.GetCommaSeparatedString(&idArr)
	data["access_token"] = token
	data["v"] = version

	return link, data
}
