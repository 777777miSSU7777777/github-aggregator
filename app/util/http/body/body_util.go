package body

import (
	"io"
	"io/ioutil"
	"net/http"
)

//ReadRequestBody returns body's bytes from request.
// Also returns error if happen.
func ReadRequestBody(req *http.Request) ([]byte, error) {
	return readBody(req.Body)
}

//ReadResponseBody returns body's bytes from response.
// Also returns error if happen.
func ReadResponseBody(resp *http.Response) ([]byte, error) {
	return readBody(resp.Body)
}

func readBody(body io.ReadCloser) ([]byte, error) {
	bodyBytes, err := ioutil.ReadAll(body)

	if err != nil {
		return nil, err
	}

	return bodyBytes, nil
}
