package telegram

// LinkFactory builds links for vk.com API methods
type LinkFactory struct{}

const apiRequestStart = "https://api.telegram.org/bot"

// BuildGetMeLink returns link for getMe method
func (factory *LinkFactory) BuildGetMeLink(token string) string {
	return apiRequestStart + token + "/getMe"
}

// BuildSendMessageLink returns link for sendMessage method
func (factory *LinkFactory) BuildSendMessageLink(token string) string {
	return apiRequestStart + token + "/sendMessage"
}

// BuildGetUpdatesLink return link for getUpdates method
func (factory *LinkFactory) BuildGetUpdatesLink(token string) string {
	return apiRequestStart + token + "/getUpdates"
}
