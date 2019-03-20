package headerutil

import (
	"net/http"
)


func ReadRequestHeader(req *http.Request) map[string][]string {
	return req.Header
}


func ReadResponseHeader(resp *http.Response) map[string][]string {
	return resp.Header
}
