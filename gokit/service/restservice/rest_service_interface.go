package restservice

import (
	"github.com/777777miSSU7777777/github-aggregator/pkg/entity"
)

type RESTService interface {
	CurrentUser() entity.User
	TokenScopes() []entity.Scope
	UserOrgs() []entity.Organization
	FilteredPulls(string, []entity.Organization) []entity.PullRequest
}
