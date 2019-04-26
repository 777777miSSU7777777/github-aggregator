package entity

// User is an entity which represents github user.
type User struct {
	Login      string `json:"login"`
	AvatarURL  string `json:"avatar_url"`
	ProfileURL string `json:"html_url"`
}
