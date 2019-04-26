package entity

// Organization is an entity which represents github organization.
type Organization struct {
	Login     string `json:"login"`
	ReposURL  string `json:"repos_url"`
	AvatarURL string `json:"avatar_url"`
}
