package cookieutil

import (
	"net/http"
	"time"
)


var expirationDuration time.Duration


// SetExpiration sets cookie duration expiration for SaveCookie func.
// If time.ParseDuration occurs any error, this will be returned.
func SetExpiration(duration string)(error){
	parsedDuration, err := time.ParseDuration(duration); if err != nil {
		return err
	}

	expirationDuration = parsedDuration

	return nil
}


// SaveCookie saves cookie with specified key and string.
func SaveCookie(rw http.ResponseWriter, key string, value string) {
	cookie := http.Cookie{Name: key, Value: value, Expires: time.Now().Add(expirationDuration)}
	http.SetCookie(rw, &cookie)
}


// GetCookieValue returns cookie value for specified key.
// Cookie value is presented as string.
// Returns error if cookie not found.
func GetCookieValue(req *http.Request, key string) (string, error) {
	cookie, err := req.Cookie(key)

	if err != nil {
		return "", err
	}

	return cookie.Value, nil
}


// DeleteCookie deletes cookie for specified key.
// Cookie key should be presented as string.
// Returns error if cookie not found.
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
