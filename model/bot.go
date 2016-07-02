package model

// Bot is a struct representing telegram bot
type Bot struct {
	accessToken string
	id          int
	username    string
}

// AccessToken returns access token of telegram bot
func (bot *Bot) AccessToken() string {
	return bot.accessToken
}

// ID return telegram bot id
func (bot *Bot) ID() int {
	return bot.id
}

// Username return telegram bot username
func (bot *Bot) Username() string {
	return bot.username
}

// Load bot info
func (bot *Bot) Load(token string, id int, username string) {
	bot.accessToken = token
	bot.id = id
	bot.username = username
}
