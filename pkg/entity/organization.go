package entity

// Organization is an entity which represents body of request to "https://api.github.com/organizations" for authenticated user.
type Organization struct {
	Login     string `json:"login"`
	ReposURL  string `json:"repos_url"`
	AvatarURL string `json:"avatar_url"`
}
