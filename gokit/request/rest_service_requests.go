package request

type CurrentUserRequest struct {
}

type TokenScopesRequest struct {
}

type UserOrgsRequest struct {
}

type FilteredPullsReq struct {
	Filter       string   `json:"filter"`
	SelectedOrgs []string `json:"selected_orgs"`
}
