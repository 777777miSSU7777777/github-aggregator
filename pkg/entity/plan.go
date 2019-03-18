package entity

type Plan struct {
	Name string `json:"name"`
	Space int `json:"space"`
	Collabarators int `json:"collabarators"`
	PrivateRepos int `json:"private_repos"`
}