package query

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/777777miSSU7777777/github-aggregator/pkg/http/headerutil"
	"github.com/777777miSSU7777777/github-aggregator/pkg/constants"
)


// Scopes returns scopes provided Github API access token.
// Acces token should be presented as string.
// Scopes is presented as string array.
// If http.Get or headerutil.ReadResponseHeader occurs any error, this will be returned.
func Scopes(tkn string)([]string, error) {
	resp, err := http.Get( fmt.Sprintf("%s%s?%s%s", constants.GHApiURL, constants.User, constants.AccessTokenParam, tkn))

	if err != nil {
		return nil, err
	}

	respHeader := headerutil.ReadResponseHeader(resp)

	return strings.Split(respHeader[constants.Scopes][0], ","), nil
}
