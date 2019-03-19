package cookieutil

import (
	"net/http"
	"time"
)


var expirationDuration time.Duration


func SetExpiration(duration string)(error){
	parsedDuration, err := time.ParseDuration(duration); if err != nil {
		return err
	}

	expirationDuration = parsedDuration

	return nil
}

func SaveCookie(rw http.ResponseWriter, key string, value string) {
	cookie := http.Cookie{Name: key, Value: value, Expires: time.Now().Add(expirationDuration)}
	http.SetCookie(rw, &cookie)
}


func GetCookieValue(req *http.Request, key string) (string, error) {
	cookie, err := req.Cookie(key)

	if err != nil {
		return "", err
	}

	return cookie.Value, nil
}


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
