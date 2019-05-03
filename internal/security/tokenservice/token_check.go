package tokenservice

import (
	"net/http"
)

const (
	// OAUTH2_HEADER key for Oauth token in request header.
	OAUTH2_HEADER = "Authorization"

	// OAUTH2_PREFIX prefix for Oauth token.
	OAUTH2_PREFIX = "Bearer "

	// GH_API_URL url to get token validity status.
	GH_API_URL = "https://api.github.com/user"

	// ACCESS_TOKEN key for access token value.
	ACCESS_TOKEN = "access_token"
)

var client *http.Client

func init() {
	client = &http.Client{}
}

type tokenChecker struct {
}

func (tc tokenChecker) checkValidity(token string) (bool, error) {
	req, err := http.NewRequest("GET", GH_API_URL, nil)

	req.Header.Set(OAUTH2_HEADER, OAUTH2_PREFIX+token)

	if err != nil {
		return false, err
	}

	resp, err := client.Do(req)

	if err != nil {
		return false, err
	}

	if resp.StatusCode == 200 {
		return true, nil
	} else if resp.StatusCode == 401 {
		return false, nil
	}

	return false, nil
}
