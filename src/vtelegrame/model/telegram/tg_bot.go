package telegram

// Bot is a struct representing telegram bot
type Bot struct {
	accessToken string
	id          int
	name        string
	username    string
	chatID      int
}

// AccessToken returns access token of telegram bot
func (bot *Bot) AccessToken() string {
	return bot.accessToken
}

// ID returns telegram bot id
func (bot *Bot) ID() int {
	return bot.id
}

// Name returns telegram bot username
func (bot *Bot) Name() string {
	return bot.name
}

func (bot *Bot) Username() string {
	return bot.username;
}

// ChatID returns id of a chat with user
func (bot *Bot) ChatID() int {
	return bot.chatID
}

// Load bot info
func (bot *Bot) Load(token string, id int, name string) {
	bot.accessToken = token
	bot.id = id
	bot.name = name
}

// SetChatID sets id of the chat with user
func (bot *Bot) SetChatID(id int) {
	bot.chatID = id
}

// SetUsername sets name of telegram user
func (bot *Bot) SetUsername(username string) {
	bot.username = username;
}
