package entity

// User is an entity which represents body of request to "https://api.github.com/" for authenticated user.
type User struct {
	Login            string `json:"login"`
	AvatarURL        string `json:"avatar_url"`
	HTMLURL          string `json:"html_url"`
	OrganizationsURL string `json:"organizations_url"`
}
