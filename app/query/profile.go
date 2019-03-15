package query

import (
	"net/http"

	"github.com/777777miSSU7777777/github-aggregator-web/app/util/http/body"
	"github.com/777777miSSU7777777/github-aggregator-web/app/util/http/cookie"
	"github.com/777777miSSU7777777/github-aggregator-web/app/util/json"
)

//GetUsername returns username from github for provided token.
// Also returns error if happen.
func GetUsername(req *http.Request) (string, error) {
	username, err := getUserField(req, "login")

	if err != nil {
		return "", err
	}

	return username.(string), nil
}

//GetAvatarURL returns avatar's url from github for provided token.
// Also returns error if happen.
func GetAvatarURL(req *http.Request) (string, error) {
	avatarURL, err := getUserField(req, "avatar_url")

	if err != nil {
		return "", err
	}

	return avatarURL.(string), nil
}

//GetProfileURL returns profile's url from github for provided token.
// Also returns error if happen.
func GetProfileURL(req *http.Request) (string, error) {
	profileURL, err := getUserField(req, "html_url")

	if err != nil {
		return "", err
	}

	return profileURL.(string), nil
}

func getUserField(req *http.Request, key string) (interface{}, error) {
	accessToken, err := cookie.GetCookieValue(req, "access_token")

	if err != nil {
		return nil, err
	}

	resp, err := http.Get("https://api.github.com/user?access_token=" + accessToken)

	if err != nil {
		return nil, err
	}

	respBody, err := body.ReadResponseBody(resp)

	if err != nil {
		return nil, err
	}

	encodedBody, err := json.BytesToMap(respBody)

	if err != nil {
		return nil, err
	}

	return encodedBody[key], nil
}
