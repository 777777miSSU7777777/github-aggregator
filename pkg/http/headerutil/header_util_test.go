package headerutil

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadHeader__SameHeader__Equal(t *testing.T) {
	header := http.Header{}

	header.Set("123", "test")

	readHeader := readHeader(header)

	assert.Equal(t, map[string][]string(header), readHeader)
}

func TestReadHeader__DifferentHeaders__NotEqual(t *testing.T) {
	header := http.Header{}

	header.Set("123", "test")

	readHeader := readHeader(header)

	header.Set("123", "test123")

	assert.NotEqual(t, map[string][]string(header), readHeader)
}

func TestRequestHeader__SameHeader__Equal(t *testing.T) {
	req := &http.Request{Header: http.Header{}}

	req.Header.Set("123", "test")

	header := ReadRequestHeader(req)

	assert.Equal(t, map[string][]string(req.Header), header)
}

func TestRequestHeader__DifferentHeaders__NotEqual(t *testing.T) {
	req := &http.Request{Header: http.Header{}}

	req.Header.Set("123", "test")

	header := ReadRequestHeader(req)

	req.Header.Set("123", "test123")

	assert.NotEqual(t, map[string][]string(req.Header), header)
}

func TestResponseHeader__SameHeader__Equal(t *testing.T) {
	resp := &http.Response{Header: http.Header{}}

	resp.Header.Set("123", "test")

	header := ReadResponseHeader(resp)

	assert.Equal(t, map[string][]string(resp.Header), header)
}

func TestResponseHeader__DifferentHeaders__NotEqual(t *testing.T) {
	resp := &http.Response{Header: http.Header{}}

	resp.Header.Set("123", "test")

	header := ReadResponseHeader(resp)

	resp.Header.Set("123", "test123")

	assert.NotEqual(t, map[string][]string(resp.Header), header)
}
