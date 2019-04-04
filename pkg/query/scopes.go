package query

// Gets scopes for provided token.

import (
	"fmt"
	"strings"

	"github.com/777777miSSU7777777/github-aggregator/pkg/constants"
	"github.com/777777miSSU7777777/github-aggregator/pkg/http/headerutil"
)

// GetScopes returns scopes provided Github API access token.
// Acces token should be presented as string.
// Scopes is presented as string array.
// If http.Get or headerutil.ReadResponseHeader occurs any error, this will be returned.
func GetScopes(tkn string) ([]string, error) {
	resp, err := httpGet(fmt.Sprintf("%s%s?%s%s", constants.GHApiURL, constants.User, constants.AccessTokenParam, tkn))

	if err != nil {
		return nil, err
	}

	respHeader := headerutil.ReadResponseHeader(resp)

	return strings.Split(respHeader[constants.Scopes][0], ","), nil
}
