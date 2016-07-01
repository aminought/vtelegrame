package app

import (
	"fmt"
	"vtelegrame/api"
	"vtelegrame/model"

	"github.com/skratchdot/open-golang/open"
)

// Application is a main struct
type Application struct {
	VKApi  *api.VKAPI
	VkUser model.VKUser
}

// Run is an entry point of application
func (application *Application) Run() {
	token := application.getAccessToken()
	application.loadVKUser(token)
}

func (application *Application) getAccessToken() string {
	fmt.Println("Hello, user. Let's authorize in vk.com.")
	var link = application.VKApi.GetAuthorizeLink()
	open.Start(link)

	fmt.Println("Please, insert here your access token:")
	var token string
	fmt.Scanf("%s", &token)

	return token
}

func (application *Application) loadVKUser(token string) {
	var fields = []string{"uid", "first_name", "last_name"}
	users := application.VKApi.GetUsers(nil, fields, "nom", token)
	if len(users) > 0 {
		application.VkUser.Load(token, users[0])
	} else {
		log.Error("Who are you?")
	}
}
