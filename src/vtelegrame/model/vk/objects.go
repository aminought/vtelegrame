package vk

// APIUser represents regular vk.com user
type APIUser struct {
	ID       int    `json:"uid"`
	Name     string `json:"first_name"`
	LastName string `json:"last_name"`
}
