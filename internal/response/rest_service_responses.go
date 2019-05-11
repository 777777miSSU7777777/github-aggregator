package response

import (
	"github.com/777777miSSU7777777/github-aggregator/pkg/entity"
)

type CurrentUserResponse struct {
	CurrentUser entity.User `json:"current_user"`
}

type TokenScopesResponse struct {
	TokenScopes []entity.Scope `json:"token_scopes"`
	Error       string         `json:"error"`
}

type UserOrgsResponse struct {
	UserOrgs []entity.Organization `json:"user_orgs"`
}

type FilteredPullsResponse struct {
	FilteredPulls []entity.PullRequest `json:"filtered_pulls"`
	Error         string               `json:"error"`
}
