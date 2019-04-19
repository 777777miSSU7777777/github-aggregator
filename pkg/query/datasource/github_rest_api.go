package datasource

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/777777miSSU7777777/github-aggregator/pkg/constants"
	"github.com/777777miSSU7777777/github-aggregator/pkg/http/bodyutil"
	"github.com/777777miSSU7777777/github-aggregator/pkg/http/headerutil"
)

// GithubRESTAPI is an implentation of data source for REST Github API v3.
type GithubRESTAPI struct {
}

// GetUser returns body of request to "https://api.github.com/user" for provided Github API access token.
// Access token should be presented as string.
// Body is presented as byte array.
// If http.Get or bodyutil.ReadResponseBody occurs any error, this will be returned.
func (ds GithubRESTAPI) GetUser(ctx context.Context, token string) ([]byte, error) {
	resp, err := http.Get(fmt.Sprintf("%s%s?%s%s", constants.GHApiURL, constants.User, constants.AccessTokenParam, token))

	if err != nil {
		return nil, err
	}

	userBody, err := bodyutil.ReadResponseBody(resp)

	if err != nil {
		return nil, err
	}

	return userBody, nil
}

// GetScopes returns scopes provided Github API access token.
// Access token should be presented as string.
// Scopes is presented as string array.
// If http.Get or headerutil.ReadResponseHeader occurs any error, this will be returned.
func (ds GithubRESTAPI) GetScopes(ctx context.Context, token string) ([]string, error) {
	resp, err := http.Get(fmt.Sprintf("%s%s?%s%s", constants.GHApiURL, constants.User, constants.AccessTokenParam, token))

	if err != nil {
		return nil, err
	}

	respHeader := headerutil.ReadResponseHeader(resp)

	return strings.Split(respHeader[constants.Scopes][0], ","), nil
}

// GetOrgs returns body of request to "https://api.github.com/organizations" for provided Github API access token.
// Access token should be presented as string.
// Body is presented as byte array.
// If http.Get or bodyutil.ReadResponseBody occurs any error, this will be returned.
func (ds GithubRESTAPI) GetOrgs(ctx context.Context, token string) ([]byte, error) {
	resp, err := http.Get(fmt.Sprintf("%s%s%s?%s%s", constants.GHApiURL, constants.User, constants.Organizations, constants.AccessTokenParam, token))

	if err != nil {
		return nil, err
	}

	orgsBody, err := bodyutil.ReadResponseBody(resp)

	if err != nil {
		return nil, err
	}

	return orgsBody, nil
}