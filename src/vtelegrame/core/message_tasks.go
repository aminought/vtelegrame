package core

import (
	tgAPI "vtelegrame/api/telegram"
	vkAPI "vtelegrame/api/vk"
	tgModel "vtelegrame/model/telegram"
	vkModel "vtelegrame/model/vk"
)

// HandleVKMessages fetch unread messages, mark it as read,
// reply with concrete message and notify with telegram bot
func HandleVKMessages(vkapi *vkAPI.API, tgapi *tgAPI.API, vkToken string,
	tgBot *tgModel.Bot, messages *Messages) {
	dialogs := getVKUnreadDialogs(vkapi, vkToken)
	unreadMessages := vkapi.GetUnreadMessages(vkToken)

	for _, v := range dialogs.Items {
		chatID := v.Message.ChatID

		crosspostMessages(v.Message, unreadMessages, vkapi, tgapi,
			vkToken, tgBot, messages)

		if chatID == 0 {
			id := v.Message.UserID
			users := vkapi.GetUsers([]int{id}, []string{"sex"}, "nom", vkToken)
			u := users[0]
			text := messages.GetVkAutoAnswer(u.Sex, tgBot.Username())
			vkapi.SendMessage(id, text, vkToken)
		}
	}
}

func crosspostMessages(lastMessage vkModel.Message, unreadMessages []vkModel.Message,
	vkapi *vkAPI.API, tgapi *tgAPI.API, vkToken string, tgBot *tgModel.Bot, messages *Messages) {

	for _, v := range unreadMessages {
		if v.UserID == lastMessage.UserID && v.ChatID == lastMessage.ChatID {
			text := v.Body
			vkapi.MarkMessageAsRead([]int{v.ID}, vkToken)
			users := vkapi.GetUsers([]int{v.UserID}, []string{"sex"}, "nom", vkToken)
			text = messages.GetTgMessageFormat(users[0].Name, users[0].LastName, v.Title, text)
			tgapi.SendMessage(tgBot, text)
			log.Info("Message was send:\n" + text)
		}
	}
}

func getVKUnreadDialogs(api *vkAPI.API, token string) vkModel.MessageItemList {
	return api.GetDialogs(0, 200, 0, true, token)
}
