package entity


// Profile is an entity for index page render with user info .
type Profile struct {
	Username string `json:"login"`
	AvatarURL string `json:"avatar_url"`
	ProfileURL string `json:"html_url"`
	Scopes []string `json:"scopes"`
}