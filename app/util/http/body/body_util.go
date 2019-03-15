package body

import (
	"net/http"
	"io/ioutil"
	"io"
)

func ReadRequestBody(req *http.Request)([]byte, error){
	return readBody(req.Body)
}

func ReadResponseBody(resp *http.Response)([]byte, error){
	return readBody(resp.Body)
}

func readBody(body io.ReadCloser)([]byte, error){
	bodyBytes, err := ioutil.ReadAll(body)

	if err != nil {
		return nil, err
	}

	return bodyBytes, nil
}