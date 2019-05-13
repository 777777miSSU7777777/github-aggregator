package rest

import (
	"github.com/777777miSSU7777777/github-aggregator/pkg/entity"
)

type currentUserResponse struct {
	CurrentUser entity.User `json:"current_user"`
}

type tokenScopesResponse struct {
	TokenScopes []entity.Scope `json:"token_scopes"`
	Error       string         `json:"error"`
}

type userOrgsResponse struct {
	UserOrgs []entity.Organization `json:"user_orgs"`
}

type filteredPullsResponse struct {
	FilteredPulls []entity.PullRequest `json:"filtered_pulls"`
	Error         string               `json:"error"`
}
