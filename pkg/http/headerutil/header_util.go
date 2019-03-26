// Package headerutil implements functions for working with heade of http request/response.
package headerutil

import (
	"net/http"
)

// ReadRequestHeader returns header of req.
// Header is presented as map of string arrays.
func ReadRequestHeader(req *http.Request) map[string][]string {
	return req.Header
}

// ReadResponseHeader returns header of resp.
// Header is presented as map of string arrays.
func ReadResponseHeader(resp *http.Response) map[string][]string {
	return resp.Header
}
