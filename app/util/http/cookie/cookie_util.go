package cookie

import(
	"net/http"
	"time"
)

func SaveCookie(rw http.ResponseWriter, key string, value string, duration time.Duration){
	cookie := http.Cookie{Name: key, Value: value, Expires: time.Now().Add(duration)}
	http.SetCookie(rw, &cookie)
}

func GetCookieValue(req *http.Request, key string)(string, error){
	cookie, err := req.Cookie(key)

	if err != nil {
		return "", err 
	}

	return cookie.Value, nil
}

func DeleteCookie(rw http.ResponseWriter, req *http.Request, key string)(error){
	cookie, err := req.Cookie(key)
	
	if err != nil {
		return err
	}

	cookie.Value = ""
	cookie.MaxAge = -1

	http.SetCookie(rw, cookie)

	return nil
}