package entity

//Session struct for index page when user is authorized with acces token
// contains auth status, username, avatar url, profile url and token scopes.
type Session struct {
	Authorized bool
	Username string
	ProfileURL string
	AvatarURL string
	Scopes []string
}