package cookie

import (
	"net/http"
	"time"
)

//SaveCookie saves value to cookie with specified key and store duration.
func SaveCookie(rw http.ResponseWriter, key string, value string, duration time.Duration) {
	cookie := http.Cookie{Name: key, Value: value, Expires: time.Now().Add(duration)}
	http.SetCookie(rw, &cookie)
}

//GetCookieValue returns value from cookie with specified key.
// Also returns error if happen.
func GetCookieValue(req *http.Request, key string) (string, error) {
	cookie, err := req.Cookie(key)

	if err != nil {
		return "", err
	}

	return cookie.Value, nil
}

//DeleteCookie removes cookie with specified key.
// Also returns error if happen.
func DeleteCookie(rw http.ResponseWriter, req *http.Request, key string) error {
	cookie, err := req.Cookie(key)

	if err != nil {
		return err
	}

	cookie.Value = ""
	cookie.MaxAge = -1

	http.SetCookie(rw, cookie)

	return nil
}
