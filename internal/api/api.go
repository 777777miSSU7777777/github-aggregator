// Package api implements functions for github aggregator rest api.
package api

import (
	"net/http"

	"github.com/go-kit/kit/log"
)

var httpGet = http.Get

var client *http.Client

var logger log.Logger

func init() {
	client = &http.Client{}
}

func SetLogger(newLogger log.Logger) {
	logger = newLogger
}
