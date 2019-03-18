package query

import (
	"net/http"
	"encoding/json"

	"github.com/777777miSSU7777777/github-aggregator-web/app/util/http/body"
	"github.com/777777miSSU7777777/github-aggregator-web/app/util/http/cookie"
	"github.com/777777miSSU7777777/github-aggregator-web/app/entity"
	"github.com/777777miSSU7777777/github-aggregator-web/app/constants"
)

func GetSessionInfo(req *http.Request)(entity.Session, error){
	session := entity.Session{}

	accessToken, err := cookie.GetCookieValue(req, "access_token")

	if err != nil {
		return session, err
	}

	session.Authorized = accessToken != ""; if session.Authorized {
		resp, err := http.Get(constants.UserApiURL + "?" + constants.AccessTokenParam + accessToken)

		if err != nil {
			return session, err
		}
	
		respBody, err := body.ReadResponseBody(resp)
	
		if err != nil {
			return session, err
		}

		err = json.Unmarshal(respBody, session)

		if err != nil {
			return session, err
		}

		return session, nil
	}

	return session, nil
}
