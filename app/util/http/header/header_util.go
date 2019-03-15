package header

import (
	"net/http"
)

//ReadRequestHeader returns header's bytes from request.
func ReadRequestHeader(req *http.Request) map[string][]string {
	return readHeader(req.Header)
}

//ReadResponseHeader returns header's bytes from response.
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
