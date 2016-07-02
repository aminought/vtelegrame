package vk

import "vtelegrame/model/vk"
import "vtelegrame/http"
import "encoding/json"
import "github.com/op/go-logging"

// API represents api for vk.com
type API struct {
	LinkFactory LinkFactory
	ClientID    string
}

var log = logging.MustGetLogger("api")

// GetAuthorizeLink returns link for authorization
func (api *API) GetAuthorizeLink() string {
	return api.LinkFactory.BuildAuthorizeLink(api.ClientID)
}

type getUsersResponse struct {
	Users []vk.APIUser `json:"response"`
}

// GetUsers returns regular vk.com users
func (api *API) GetUsers(ids []int, fields []string, nameCase string, token string) []vk.APIUser {
	link := api.LinkFactory.BuildGetUsersLink(ids, fields, nameCase, token)
	data := http.GetRequest(link)
	response := getUsersResponse{}
	json.Unmarshal(data, &response)

	return response.Users
}
