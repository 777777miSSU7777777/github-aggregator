package entity

type Profile struct {
	Username string `json:"login"`
	AvatarURL string `json:"avatar_url"`
	ProfileURL string `json:"html_url"`
	Scopes []string `json:"scopes"`
}