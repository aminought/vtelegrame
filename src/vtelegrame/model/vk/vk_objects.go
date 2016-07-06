package vk

// APIUser represents regular vk.com user
type APIUser struct {
	ID       int    `json:"id"`
	Name     string `json:"first_name"`
	LastName string `json:"last_name"`
	Sex      int    `json:"sex"`
}

// Message represents vk.com message
type Message struct {
	ID        int    `json:"id"`
	Date      int    `json:"date"`
	UserID    int    `json:"user_id"`
	ReadState int    `json:"read_state"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	ChatID    int    `json:"chat_id"`
}

// MessageItem represents item of a MessageList
type MessageItem struct {
	Unread  int     `json:"unread"`
	Message Message `json:"message"`
	InRead  int     `json:"in_read"`
	OutRead int     `json:"out_read"`
}

// MessageItemList represents vk.com list of message items
type MessageItemList struct {
	Count int           `json:"count"`
	Items []MessageItem `json:"items"`
}

// MessageList represents vk.com list of messages
type MessageList struct {
	Count    int       `json:"count"`
	Messages []Message `json:"items"`
}
