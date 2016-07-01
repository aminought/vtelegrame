package model

// VKAPIUser represents regular vk.com user
type VKAPIUser struct {
	ID       int    `json:"uid"`
	Name     string `json:"first_name"`
	LastName string `json:"last_name"`
}
