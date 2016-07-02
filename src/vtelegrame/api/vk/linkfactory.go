package vk

// LinkFactory is a link factory for vk.com API
type LinkFactory struct{}

const apiRequestStart = "https://api.vk.com/method/"

// BuildAuthorizeLink returns link for authorization
func (factory *LinkFactory) BuildAuthorizeLink(clientID string) string {
	var link = "https://oauth.vk.com/authorize?"
	link += "client_id=" + clientID + "&"
	link += "scope=friends,status,wall,groups,messages,notifications,offline&"
	link += "redirect_uri=http://oauth.vk.com/blank.html&"
	link += "display=page&"
	link += "response_type=token&"
	link += "v=5.52"
	return link
}

// BuildGetUsersLink return link for users.get method
func (factory *LinkFactory) BuildGetUsersLink(ids []int, fields []string, nameCase string, token string) string {
	const method = "users.get"

	var link = apiRequestStart
	link += method + "?"
	link += "fields="
	var fieldsLen = len(fields)
	for i, v := range fields {
		link += v
		if i < fieldsLen-1 {
			link += ","
		}
	}
	link += "&"
	link += "name_case=" + nameCase + "&"
	link += "access_token=" + token

	return link
}
