package core

import (
	"encoding/json"
	"io/ioutil"
	"strings"
)

// Messages represents messages from messages.json
type Messages struct {
	VkAutoAnswer       []string `json:"vk_auto_answer"`
	VkAutoAnswerShorts []string `json:"vk_auto_answer_shorts"`
	TgStartMessage     string   `json:"tg_start_message"`
	TgMessageFormat    []string `json:"tg_message_format"`
}

const (
	tglogin    = "%tg_login%"
	firstname  = "%firstname%"
	lastname   = "%lastname%"
	sex        = "%sex%"
	sex2       = "%sex2%"
	chatbraces = "%chat_braces%"
	text       = "%text%"
)

const (
	female  = iota
	male    = iota
	unknown = iota
)

type pair struct {
	keyword string
	word    string
}

// Load messages
func (messages *Messages) Load(path string) {
	data, _ := ioutil.ReadFile(path)
	json.Unmarshal(data, messages)
}

// GetVkAutoAnswer returns the main answer to vk.com message
func (messages *Messages) GetVkAutoAnswer(sexP int, tgloginP string) string {
	answer := strings.Join(messages.VkAutoAnswer, "\n")
	return messages.replaceKeywords(answer, pair{sex, messages.getSex(sexP)}, pair{tglogin, tgloginP})
}

// GetTgMessageFormat returns message for crossposting to telegram
func (messages *Messages) GetTgMessageFormat(firstnameP string, lastnameP string,
	chatTitle string, textP string) string {
	answer := strings.Join(messages.TgMessageFormat, "\n")
	chat := messages.getChatBraces(chatTitle)
	return messages.replaceKeywords(answer, pair{firstname, firstnameP},
		pair{lastname, lastnameP}, pair{chatbraces, chat}, pair{text, textP})
}

// GetTgStartMessageFormat returns start message in telegram
func (messages *Messages) GetTgStartMessageFormat(tgLoginP string) string {
	return messages.replaceKeywords(messages.TgStartMessage, pair{tglogin, tgLoginP})
}

func (messages *Messages) getSex(sexP int) string {
	if sexP == male {
		return ""
	}
	return "а"
}

func (messages *Messages) getChatBraces(chatTitle string) string {
	chat := ""
	if chatTitle != " ... " {
		chat = "(" + chatTitle + ")"
	}
	return chat
}

func (messages *Messages) replaceKeywords(message string, pairs ...pair) string {
	text := message
	for _, v := range pairs {
		text = messages.replaceKeyword(text, v.keyword, v.word)
	}
	return text
}

func (messages *Messages) replaceKeyword(text string, keyword string, word string) string {
	return strings.Replace(text, keyword, word, -1)
}
