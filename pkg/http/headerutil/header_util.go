// Package headerutil implements functions for working with heade of http request/response.
package headerutil

import (
	"net/http"
)

func readHeader(header http.Header) map[string][]string {
	hdr := map[string][]string{}

	for k, v := range header {
		hdr[k] = v
	}

	return hdr
}

// ReadRequestHeader returns header of req.
func ReadRequestHeader(req *http.Request) map[string][]string {
	return readHeader(req.Header)
}

// ReadResponseHeader returns header of resp.
func ReadResponseHeader(resp *http.Response) map[string][]string {
	return readHeader(resp.Header)
}
