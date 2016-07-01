package api

import "vtelegrame/model"
import "vtelegrame/http"
import "encoding/json"
import "github.com/op/go-logging"

// VKAPI is a class representing api for vk.com
type VKAPI struct {
	ClientID string
}

var log = logging.MustGetLogger("api")

const apiRequestStart = "https://api.vk.com/method/"

// GetAuthorizeLink returns link for authorization
func (vkAPI *VKAPI) GetAuthorizeLink() string {
	var link = "https://oauth.vk.com/authorize?"
	link += "client_id=" + vkAPI.ClientID + "&"
	link += "scope=friends,status,wall,groups,messages,notifications,offline&"
	link += "redirect_uri=http://oauth.vk.com/blank.html&"
	link += "display=page&"
	link += "response_type=token&"
	link += "v=5.52"
	return link
}

func (vkAPI *VKAPI) getGetUsersLink(ids []int, fields []string, nameCase string, token string) string {
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

type getUsersResponse struct {
	Users []model.VKAPIUser `json:"response"`
}

// GetUsers returns regular vk.com users
func (vkAPI *VKAPI) GetUsers(ids []int, fields []string, nameCase string, token string) []model.VKAPIUser {
	link := vkAPI.getGetUsersLink(ids, fields, nameCase, token)
	data := http.Request(link)
	response := getUsersResponse{}
	json.Unmarshal(data, &response)

	return response.Users
}
