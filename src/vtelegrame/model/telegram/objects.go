package telegram

// User represents telegram user
type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserName  string `json:"username"`
}

// Chat represents chat object
type Chat struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserName  string `json:"username"`
	Type      string `json:"type"`
	Title     string `json:"title"`
}

// Message represents message object
type Message struct {
	ID   int    `json:"message_id"`
	From User   `json:"from"`
	Date int    `json:"date"`
	Chat Chat   `json:"chat"`
	Text string `json:"text"`
}

// Update represets update object
type Update struct {
	ID      int     `json:"update_id"`
	Message Message `json:"message"`
}
