package datasrcfactory

import (
	"github.com/777777miSSU7777777/github-aggregator/pkg/query/datasource"
)

func New(dataSrcType string) datasource.DataSource {
	switch dataSrcType {
	case "rest-api":
		return &datasource.GithubRESTAPI{}
	default:
		return &datasource.GithubRESTAPI{}
	}
}
