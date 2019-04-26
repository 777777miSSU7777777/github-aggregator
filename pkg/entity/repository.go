package entity

// Owner alias to user entity.
type Owner User

// Repository is entity which represents github repository.
type Repository struct {
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	Owner    Owner  `json:"owner"`
	HTMLURL  string `json:"html_url"`
	PullsURL string `json:"pulls_url"`
}
