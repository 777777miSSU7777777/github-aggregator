package query

import (
	"net/http"
	"github-aggregator-web/app/util/http/cookie"
	"github-aggregator-web/app/util/http/body"
	"github-aggregator-web/app/util/json"
)

func GetUsername(req *http.Request)(string, error){
	username, err := getUserField(req, "login")

	if err != nil {
		return "", err
	}

	return username.(string), nil
}

func GetAvatarUrl(req *http.Request)(string, error){
	avatarUrl, err := getUserField(req, "avatar_url")

	if err != nil {
		return "", err
	}

	return avatarUrl.(string), nil
}

func GetProfileUrl(req *http.Request)(string, error){
	profileUrl, err := getUserField(req, "html_url")

	if err != nil {
		return "", err
	}

	return profileUrl.(string), nil 
}

func getUserField(req *http.Request, key string)(interface{}, error){
	access_token, err := cookie.GetCookieValue(req, "access_token")

	if err != nil {
		return nil, err
	}

	resp, err := http.Get("https://api.github.com/user?access_token=" + access_token)

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