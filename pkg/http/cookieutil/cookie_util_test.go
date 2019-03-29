package cookieutil

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSetDurationExpiration__SameDuration__Equal(t *testing.T) {
	testDuration, _ := time.ParseDuration("1h")

	SetExpiration("1h")

	assert.Equal(t, testDuration, expirationDuration)
}

func TestSetDurationExpiration__DifferentDuration__NotEqual(t *testing.T) {
	testDuration, _ := time.ParseDuration("30m")

	SetExpiration("1h")

	assert.NotEqual(t, testDuration, expirationDuration)
}

func TestSetDurationExpiration__IncorrectDurationString__Error(t *testing.T) {
	err := SetExpiration("")

	assert.Error(t, err)
}

func TestSaveCookie__SameValue__Equal(t *testing.T) {
	rw := httptest.NewRecorder()

	SaveCookie(rw, "123", "test")

	req := &http.Request{Header: http.Header{"Cookie": rw.HeaderMap["Set-Cookie"]}}

	cookie, _ := req.Cookie("123")

	assert.Equal(t, "test", cookie.Value)
}

func TestSaveCookie__DifferentValues__NotEqual(t *testing.T) {
	rw := httptest.NewRecorder()

	SaveCookie(rw, "123", "test")

	req := &http.Request{Header: http.Header{"Cookie": rw.HeaderMap["Set-Cookie"]}}

	cookie, _ := req.Cookie("123")

	assert.NotEqual(t, "123", cookie.Value)
}

func TestGetCookieValue__SameValue__Equal(t *testing.T) {
	rw := httptest.NewRecorder()

	SaveCookie(rw, "123", "test")

	req := &http.Request{Header: http.Header{"Cookie": rw.HeaderMap["Set-Cookie"]}}

	cookieValue, _ := GetCookieValue(req, "123")

	assert.Equal(t, "test", cookieValue)
}

func TestGetCookieValue__DifferentValues__NotEqual(t *testing.T) {
	rw := httptest.NewRecorder()

	SaveCookie(rw, "123", "test")

	req := &http.Request{Header: http.Header{"Cookie": rw.HeaderMap["Set-Cookie"]}}

	cookieValue, _ := GetCookieValue(req, "123")

	assert.NotEqual(t, "123", cookieValue)
}

func TestGetCookieValue__CookieNotFound__Error(t *testing.T) {
	req := &http.Request{}

	_, err := GetCookieValue(req, "123")

	assert.Error(t, err)
}

func TestDeleteCookie__ExistingCookie__Successful(t *testing.T) {
	rw := httptest.NewRecorder()

	SaveCookie(rw, "123", "test")

	req := &http.Request{Header: http.Header{"Cookie": rw.HeaderMap["Set-Cookie"]}}

	err := DeleteCookie(rw, req, "123")

	assert.Nil(t, err)
}

func TestDeleteCookie__CookieNotFound__Error(t *testing.T) {
	rw := httptest.NewRecorder()

	req := &http.Request{}

	err := DeleteCookie(rw, req, "123")

	assert.Error(t, err)
}
