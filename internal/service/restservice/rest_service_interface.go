package restservice

import (
	"github.com/777777miSSU7777777/github-aggregator/pkg/entity"
)

type RESTService interface {
	CurrentUser() entity.User
	TokenScopes() ([]entity.Scope, error)
	UserOrgs() []entity.Organization
	FilteredPulls(string, []string) ([]entity.PullRequest, error)
}
