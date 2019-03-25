package bodyutil

import (
	"io"
	"io/ioutil"
	"net/http"
)

// ReadRequestBody returns body of req.
// Body is presented as byte array.
// Returns error if any occurs during the body reading.
func ReadRequestBody(req *http.Request) ([]byte, error) {
	return readBody(req.Body)
}

// ReadResponseBody returns body of req.
// Body is presented as byte array.
// Returns error if any occurs during the body reading.
func ReadResponseBody(resp *http.Response) ([]byte, error) {
	return readBody(resp.Body)
}

func readBody(body io.ReadCloser) ([]byte, error) {
	defer body.Close()

	bodyBytes, err := ioutil.ReadAll(body)

	if err != nil {
		return nil, err
	}

	return bodyBytes, nil
}
