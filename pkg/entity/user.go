package entity

// User is an entity for navbar with user info.
type User struct {
	Login      string `json:"login"`
	AvatarURL  string `json:"avatar_url"`
	ProfileURL string `json:"html_url"`
}
