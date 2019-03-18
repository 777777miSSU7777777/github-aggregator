package entity


type Session struct {
	Authorized bool
	Username   string
	ProfileURL string
	AvatarURL  string
	Scopes     []string
}
