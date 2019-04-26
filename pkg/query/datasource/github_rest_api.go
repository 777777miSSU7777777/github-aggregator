package datasource

import (
	"context"
	"net/http"
	"strings"

	"github.com/777777miSSU7777777/github-aggregator/pkg/entity"
	"github.com/777777miSSU7777777/github-aggregator/pkg/http/bodyutil"
	"github.com/777777miSSU7777777/github-aggregator/pkg/http/headerutil"
)

// GithubRESTAPI is an implentation of data source for REST Github API v3.
type GithubRESTAPI struct {
	client *http.Client
}

// NewGithubRESTAPI constructor for GithubRESTAPI struct.
func NewGithubRESTAPI() *GithubRESTAPI {
	return &GithubRESTAPI{client: &http.Client{}}
}

const (
	// OAUTH2_HEADER key for Oauth token in request header.
	OAUTH2_HEADER = "Authorization"

	// OAUTH2_PREFIX Prefix for Oauth token.
	OAUTH2_PREFIX = "Bearer "

	// SCOPES_HEADER key for github oauth scopes in request header.
	SCOPES_HEADER = "X-Oauth-Scopes"

	// USER_QUERY api url for current user.
	USER_QUERY = "https://api.github.com/user"

	// ORGS_QUERY api url for orgs of current user.
	ORGS_QUERY = "https://api.github.com/user/orgs"
)

// GetUser returns body of request to "https://api.github.com/user" for provided Github API access token.
// Access token should be presented as string.
// Body is presented as byte array.
// If http.Get or bodyutil.ReadResponseBody occurs any error, this will be returned.
func (ds GithubRESTAPI) GetUser(ctx context.Context, token string) ([]byte, error) {
	req, err := http.NewRequest("GET", USER_QUERY, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set(OAUTH2_HEADER, OAUTH2_PREFIX+token)

	resp, err := ds.client.Do(req)

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
	req, err := http.NewRequest("GET", USER_QUERY, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set(OAUTH2_HEADER, OAUTH2_PREFIX+token)

	resp, err := ds.client.Do(req)

	if err != nil {
		return nil, err
	}

	respHeader := headerutil.ReadResponseHeader(resp)

	return strings.Split(respHeader[SCOPES_HEADER][0], ","), nil
}

// GetOrgs returns body of request to "https://api.github.com/organizations" for provided Github API access token.
// Access token should be presented as string.
// Body is presented as byte array.
// If http.Get or bodyutil.ReadResponseBody occurs any error, this will be returned.
func (ds GithubRESTAPI) GetOrgs(ctx context.Context, token string) ([]byte, error) {
	req, err := http.NewRequest("GET", ORGS_QUERY, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set(OAUTH2_HEADER, OAUTH2_PREFIX+token)

	resp, err := ds.client.Do(req)

	if err != nil {
		return nil, err
	}

	orgsBody, err := bodyutil.ReadResponseBody(resp)

	if err != nil {
		return nil, err
	}

	return orgsBody, nil
}

// GetOrgsRepos returns body of request to "https://api.github.com/orgs/:org/repos" for provided Github API access token.
// Access token should be presented as string.
// Body is presented as byte array.
// If http.Get or bodyutil.ReadResponseBody occurs any error, this will be returned.
func (ds GithubRESTAPI) GetOrgsRepos(ctx context.Context, token string, orgs []entity.Organization) ([][]byte, error) {
	resultSetBytes := [][]byte{}

	for _, org := range orgs {
		req, err := http.NewRequest("GET", org.ReposURL, nil)

		if err != nil {
			return nil, err
		}

		req.Header.Set(OAUTH2_HEADER, OAUTH2_PREFIX+token)

		resp, err := ds.client.Do(req)

		if err != nil {
			return nil, err
		}

		reposBody, err := bodyutil.ReadResponseBody(resp)

		if err != nil {
			return nil, err
		}

		resultSetBytes = append(resultSetBytes, reposBody)
	}

	return resultSetBytes, nil
}

// GetOrgsPullRequests blank
func (ds GithubRESTAPI) GetReposPullRequests(ctx context.Context, token string, repos []entity.Repository) ([][]byte, error) {
	resultSetBytes := [][]byte{}

	for _, repo := range repos {
		req, err := http.NewRequest("GET", trimPullsURL(repo.PullsURL), nil)

		if err != nil {
			return nil, err
		}

		req.Header.Set(OAUTH2_HEADER, OAUTH2_PREFIX+token)

		resp, err := ds.client.Do(req)

		if err != nil {
			return nil, err
		}

		pullsBody, err := bodyutil.ReadResponseBody(resp)

		if err != nil {
			return nil, err
		}

		resultSetBytes = append(resultSetBytes, pullsBody)
	}

	return resultSetBytes, nil
}

func trimPullsURL(pullsUrls string) string {
	return strings.Trim(pullsUrls, "{/number/}")
}
