package datasource

import (
	"context"
)

// DataSource is an interface for github apis (REST, GRAPH-QL).
type DataSource interface {
	GetUser(context.Context, string) ([]byte, error)
	GetScopes(context.Context, string) ([]string, error)
	GetOrgs(context.Context, string) ([]byte, error)
}
