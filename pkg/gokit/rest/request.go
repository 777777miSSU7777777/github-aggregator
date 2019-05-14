package rest

type currentUserRequest struct {
}

type tokenScopesRequest struct {
}

type userOrgsRequest struct {
}

type filteredPullsReq struct {
	Filter       string   `json:"filter"`
	SelectedOrgs []string `json:"selected_orgs"`
}
