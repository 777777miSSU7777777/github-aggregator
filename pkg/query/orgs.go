package query

// Gets organizations info for provided token, if it exists.

import (
	"fmt"

	"github.com/777777miSSU7777777/github-aggregator/pkg/constants"
)

// GetOrganizations returns body of request to "https://api.github.com/organizations" for provided Github API access token.
// Access token should be presented as string.
// Body is presented as byte array.
// If http.Get or bodyutil.ReadResponseBody occurs any error, this will be returned.
func GetOrgs(tkn string) ([]byte, error) {
	resp, err := httpGet(fmt.Sprintf("%s%s%s?%s%s", constants.GHApiURL, constants.User, constants.Organizations, constants.AccessTokenParam, tkn))

	if err != nil {
		return nil, err
	}

	orgsBody, err := readResponseBody(resp)

	if err != nil {
		return nil, err
	}

	return orgsBody, nil
}
