package telegram

import "strconv"

// LinkFactory builds links for vk.com API methods
type LinkFactory struct{}

const apiRequestStart = "https://api.telegram.org/bot"

// BuildGetMeLink returns link for getMe method
func (factory *LinkFactory) BuildGetMeLink(token string) string {
	return apiRequestStart + token + "/getMe"
}

// BuildSendMessageLink returns link for sendMessage method
func (factory *LinkFactory) BuildSendMessageLink(chatID int, text string,
	token string) (string, map[string]string) {

	link := apiRequestStart + token + "/sendMessage"
	data := make(map[string]string)
	data["chat_id"] = strconv.Itoa(chatID)
	data["text"] = text
	return link, data
}

// BuildGetUpdatesLink return link for getUpdates method
func (factory *LinkFactory) BuildGetUpdatesLink(token string) string {
	return apiRequestStart + token + "/getUpdates"
}
