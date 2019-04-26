// Package api implements functions for github aggregator rest api.
package api

import (
	"net/http"
)

var httpGet = http.Get

var client *http.Client

func init() {
	client = &http.Client{}
}
