package vk

import (
	"encoding/json"
	"vtelegrame/http"
	"vtelegrame/model/vk"

	"github.com/op/go-logging"
)

// API represents api for vk.com
type API struct {
	LinkFactory LinkFactory
	ClientID    string
}

var log = logging.MustGetLogger("api")

// GetAuthorizeLink returns link for authorization
func (api *API) GetAuthorizeLink() (string, map[string]string) {
	return api.LinkFactory.BuildAuthorizeLink(api.ClientID)
}

type getUsersResponse struct {
	Users []vk.APIUser `json:"response"`
}

// GetUsers returns regular vk.com users
func (api *API) GetUsers(ids []int, fields []string, nameCase string, token string) []vk.APIUser {
	link, data := api.LinkFactory.BuildGetUsersLink(ids, fields, nameCase, token)
	answer := http.GetRequest(link, data)
	response := getUsersResponse{}
	json.Unmarshal(answer, &response)

	return response.Users
}

type getDialogsResponse struct {
	Count           int                `json:"count"`
	MessageItemList vk.MessageItemList `json:"response"`
}

// GetDialogs return vk.com dialogs
func (api *API) GetDialogs(offset int, count int, previewLength int,
	unread bool, token string) vk.MessageItemList {

	link, data := api.LinkFactory.BuildGetDialogsLink(offset, count, previewLength, unread, token)
	answer := http.GetRequest(link, data)
	response := getDialogsResponse{}
	json.Unmarshal(answer, &response)

	return response.MessageItemList
}

// SendMessage sends message to another user
func (api *API) SendMessage(userID int, message string, token string) {
	link, data := api.LinkFactory.BuildSendMessageLink(userID, message, token)
	http.PostRequest(link, data)
}

type getMessagesResponse struct {
	MessageList vk.MessageList `json:"response"`
}

// GetMessages returns vk.com messages
func (api *API) GetMessages(offset int, count int, filters int,
	previewLength int, token string) []vk.Message {
	link, data := api.LinkFactory.BuildGetMessagesLink(offset, count, filters, previewLength, token)
	answer := http.GetRequest(link, data)
	response := getMessagesResponse{}
	json.Unmarshal(answer, &response)

	return response.MessageList.Messages
}

// GetUnreadMessages returns vk.com unread messages
func (api *API) GetUnreadMessages(token string) []vk.Message {
	messages := api.GetMessages(0, 200, 0, 0, token)
	var unread []vk.Message
	for i := len(messages) - 1; i >= 0; i-- {
		if messages[i].ReadState == 0 {
			unread = append(unread, messages[i])
		}
	}
	return unread
}

// GetMessagesByID returs vk.com messages by id
func (api *API) GetMessagesByID(ids []int, previewLength int, token string) []vk.Message {
	link, data := api.LinkFactory.BuildGetMessagesByIDLink(ids, previewLength, token)
	answer := http.GetRequest(link, data)
	response := getMessagesResponse{}
	json.Unmarshal(answer, &response)

	return response.MessageList.Messages
}

// MarkMessageAsRead marks vk.com messages as read
func (api *API) MarkMessageAsRead(ids []int, token string) {
	link, data := api.LinkFactory.BuildMarkMessageAsReadLink(ids, token)
	http.GetRequest(link, data)
}
