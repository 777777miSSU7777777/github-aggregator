package entity


// Plan is an entity which represents User's entity field for Github paid plan.
type Plan struct {
	Name string `json:"name"`
	Space int `json:"space"`
	Collabarators int `json:"collabarators"`
	PrivateRepos int `json:"private_repos"`
}