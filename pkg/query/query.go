// Package query implements functions for working with Github API.
package query

import (
	"github.com/777777miSSU7777777/github-aggregator/pkg/query/datasource"
)

var dataSrc datasource.DataSource

func SetDataSource(dataSource datasource.DataSource) {
	dataSrc = dataSource
}

func GetDataSource() datasource.DataSource {
	return dataSrc
}
