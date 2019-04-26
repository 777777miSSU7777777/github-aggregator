package datasrcfactory

import (
	"github.com/777777miSSU7777777/github-aggregator/pkg/query/datasource"
)

// New returns new instance of DataSource.
func New(dataSrcType string) datasource.DataSource {
	switch dataSrcType {
	case "rest-api":
		return datasource.NewGithubRESTAPI()
	default:
		return datasource.NewGithubRESTAPI()
	}
}
