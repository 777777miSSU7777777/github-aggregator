// Package query implements functions for working with Github API.
package query

import (
	"net/http"

	"github.com/777777miSSU7777777/github-aggregator/pkg/http/bodyutil"
)

var httpGet = http.Get

var readResponseBody = bodyutil.ReadResponseBody
