package core

import (
	tgAPI "vtelegrame/api/telegram"
	vkAPI "vtelegrame/api/vk"
	tgModel "vtelegrame/model/telegram"
	vkModel "vtelegrame/model/vk"
)

// HandleVKMessages fetch unread messages, mark it as read,
// reply with concrete message and notify with telegram bot
func HandleVKMessages(vkapi *vkAPI.API, tgapi *tgAPI.API, vkToken string, tgBot *tgModel.Bot) {
	dialogs := getVKUnreadDialogs(vkapi, vkToken)
	unreadMessages := vkapi.GetUnreadMessages(vkToken)

	for _, v := range dialogs.Items {
		chatID := v.Message.ChatID

		crosspostMessages(v.Message, unreadMessages, vkapi, tgapi, vkToken, tgBot)

		if chatID == 0 {
			id := v.Message.UserID
			vkapi.SendMessage(id, "Привет. Я теперь тут: https://telegram.me/aminought. Присоединяйся!", vkToken)
		}
	}
}

func crosspostMessages(lastMessage vkModel.Message, unreadMessages []vkModel.Message,
	vkapi *vkAPI.API, tgapi *tgAPI.API, vkToken string, tgBot *tgModel.Bot) {

	for _, v := range unreadMessages {
		if v.UserID == lastMessage.UserID && v.ChatID == lastMessage.ChatID {
			text := v.Body
			vkapi.MarkMessageAsRead([]int{v.ID}, vkToken)
			tgapi.SendMessage(tgBot, text)
			log.Info("Message \"" + text + "\" was sent")
		}
	}
}

func getVKUnreadDialogs(api *vkAPI.API, token string) vkModel.MessageItemList {
	return api.GetDialogs(0, 200, 0, true, token)
}
