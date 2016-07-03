package vk

import "vtelegrame/util"
import "strconv"

// LinkFactory is a link factory for vk.com API
type LinkFactory struct{}

const apiRequestStart = "https://api.vk.com/method/"
const version = "5.52"

// BuildAuthorizeLink returns link for authorization
func (factory *LinkFactory) BuildAuthorizeLink(clientID string) string {
	var link = "https://oauth.vk.com/authorize?"
	link += "client_id=" + clientID
	link += "&scope=friends,status,wall,groups,messages,notifications,offline"
	link += "&redirect_uri=http://oauth.vk.com/blank.html"
	link += "&display=page"
	link += "&response_type=token"
	link += "&v=" + version
	return link
}

// BuildGetUsersLink returns link for users.get method
func (factory *LinkFactory) BuildGetUsersLink(ids []int, fields []string, nameCase string, token string) string {
	const method = "users.get"

	idArr := util.IntArray(ids)
	fieldArr := util.StringArray(fields)

	link := apiRequestStart + method + "?"
	link += "ids=" + util.GetCommaSeparatedString(&idArr)
	link += "&fields=" + util.GetCommaSeparatedString(&fieldArr)
	link += "&name_case=" + nameCase
	link += "&access_token=" + token
	link += "&v=" + version

	return link
}

// BuildGetDialogsLink returns link for messages.getDialogs method
func (factory *LinkFactory) BuildGetDialogsLink(offset int, count int,
	previewLength int, unread bool, token string) string {

	const method = "messages.getDialogs"
	link := apiRequestStart + method + "?"
	link += "offset=" + strconv.Itoa(offset)
	link += "&count=" + strconv.Itoa(count)
	link += "&preview_length=" + strconv.Itoa(previewLength)
	u := "1"
	if !unread {
		u = "0"
	}
	link += "&unread=" + u
	link += "&access_token=" + token
	link += "&v=" + version

	return link
}

// BuildSendMessageLink returns link for mesages.send method
func (factory *LinkFactory) BuildSendMessageLink() string {
	const method = "messages.send"
	link := apiRequestStart + method
	return link
}

//BuildGetMessagesLink returns link for messages.get method
func (factory *LinkFactory) BuildGetMessagesLink(offset int,
	count int, filters int, previewLength int, token string) string {

	const method = "messages.get"
	link := apiRequestStart + method + "?"
	link += "offset=" + strconv.Itoa(offset)
	link += "&count=" + strconv.Itoa(count)
	link += "&filters=" + strconv.Itoa(filters)
	link += "&preview_length=" + strconv.Itoa(previewLength)
	link += "&access_token=" + token
	link += "&v=" + version

	return link
}

// BuildGetMessagesByIDLink returns link for messages.getById method
func (factory *LinkFactory) BuildGetMessagesByIDLink(ids []int,
	previewLength int, token string) string {

	const method = "messages.getById"
	link := apiRequestStart + method + "?"
	idArr := util.IntArray(ids)
	link += "message_ids=" + util.GetCommaSeparatedString(&idArr)
	link += "&preview_length=" + strconv.Itoa(previewLength)
	link += "&access_token=" + token
	link += "&v=" + version

	return link
}

// BuildMarkMessageAsRead returns link for messages.markAsRead method
func (factory *LinkFactory) BuildMarkMessageAsReadLink(ids []int, token string) string {
	const method = "messages.markAsRead"
	link := apiRequestStart + method + "?"
	idArr := util.IntArray(ids)
	link += "message_ids=" + util.GetCommaSeparatedString(&idArr)
	link += "&access_token=" + token
	link += "&v=" + version

	return link
}
