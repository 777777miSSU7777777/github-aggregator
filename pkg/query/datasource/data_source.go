package datasource

import (
	"context"

	"github.com/777777miSSU7777777/github-aggregator/pkg/entity"
)

// DataSource is an interface for github apis (REST, GRAPH-QL).
type DataSource interface {
	GetUser(context.Context, string) ([]byte, error)
	GetScopes(context.Context, string) ([]string, error)
	GetOrgs(context.Context, string) ([]byte, error)
	GetOrgsRepos(context.Context, string, []entity.Organization) ([][]byte, error)
	GetReposPullRequests(context.Context, string, []entity.Repository) ([][]byte, error)
}
