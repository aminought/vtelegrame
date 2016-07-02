package telegram

// LinkFactory builds links for vk.com API methods
type LinkFactory struct{}

const apiRequestStart = "https://api.telegram.org/bot"

// BuildGetMeLink returns link for getMe method
func (factory *LinkFactory) BuildGetMeLink(token string) string {
	link := apiRequestStart + token + "/getMe"
	return link
}

// BuildSendMessageLink returns link for sendMessage method
func (factory *LinkFactory) BuildSendMessageLink(token string) string {
	link := apiRequestStart + token + "/sendMessage"
	return link
}

// BuildGetUpdatesLink return link for getUpdates method
func (factory *LinkFactory) BuildGetUpdatesLink(token string) string {
	link := apiRequestStart + token + "/getUpdates"
	return link
}
