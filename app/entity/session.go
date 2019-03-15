package entity

type Session struct {
	Authorized bool
	Username string
	ProfileUrl string
	AvatarUrl string
	Scopes []string
}