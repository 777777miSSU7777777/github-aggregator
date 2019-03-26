package query

// Gets user info for provided token, if it exists.

import (
	"fmt"
	"net/http"

	"github.com/777777miSSU7777777/github-aggregator/pkg/constants"
	"github.com/777777miSSU7777777/github-aggregator/pkg/http/bodyutil"
)

// GetUser returns body of request to "https://api.github.com/user" for provided Github API access token.
// Access token should be presented as string.
// Body is presented as byte array.
// If http.Get or bodyutil.ReadResponseBody occurs any error, this will be returned.
func GetUser(tkn string) ([]byte, error) {
	resp, err := http.Get(fmt.Sprintf("%s%s?%s%s", constants.GHApiURL, constants.User, constants.AccessTokenParam, tkn))

	if err != nil {
		return nil, err
	}

	userBody, err := bodyutil.ReadResponseBody(resp)

	if err != nil {
		return nil, err
	}

	return userBody, nil
}
