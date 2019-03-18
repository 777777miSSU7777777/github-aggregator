package headerutil

import (
	"net/http"
)


func ReadRequestHeader(req *http.Request) map[string][]string {
	return readHeader(req.Header)
}


func ReadResponseHeader(resp *http.Response) map[string][]string {
	return readHeader(resp.Header)
}


func readHeader(hdr http.Header) map[string][]string {
	header := map[string][]string{}

	for key, value := range hdr {
		header[key] = value
	}

	return hdr
}
