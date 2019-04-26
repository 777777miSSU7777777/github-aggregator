// Package bodyutil implements functions for working with body of http request/response.
package bodyutil

import (
	"io"
	"io/ioutil"
	"net/http"
)

var ioutilReadAll = ioutil.ReadAll

// ReadRequestBody returns body of req.
func ReadRequestBody(req *http.Request) ([]byte, error) {
	return readBody(req.Body)
}

// ReadResponseBody returns body of req.
func ReadResponseBody(resp *http.Response) ([]byte, error) {
	return readBody(resp.Body)
}

func readBody(body io.ReadCloser) ([]byte, error) {
	defer body.Close()

	bodyBytes, err := ioutilReadAll(body)

	if err != nil {
		return nil, err
	}

	return bodyBytes, nil
}
