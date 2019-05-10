package datasource

import (
	"context"
	"net/http"
	"strings"
	"sync"

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

// GetUser returns body with user data.
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

// GetScopes returns scopes for provided token.
func (ds GithubRESTAPI) GetScopes(ctx context.Context, token string) ([]entity.Scope, error) {
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

	return []entity.Scope(strings.Split(respHeader[SCOPES_HEADER][0], ",")), nil
}

// GetOrgs returns body with user orgs.
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

// GetOrgsRepos returns body with orgs repos.
func (ds GithubRESTAPI) GetOrgsRepos(ctx context.Context, token string, orgs []entity.Organization) ([][]byte, error) {
	resultSetBytes := [][]byte{}

	orgChan := make(chan entity.Organization, len(orgs))
	reposBytesChan := make(chan []byte, 1)
	errorsChan := make(chan error, 1)

	for _, org := range orgs {
		orgChan <- org

		go ds.getOrgRepos(ctx, token, orgChan, reposBytesChan, errorsChan)

		reposBody := <-reposBytesChan

		err := <-errorsChan

		if err != nil {
			return nil, err
		}

		resultSetBytes = append(resultSetBytes, reposBody)
	}

	return resultSetBytes, nil
}

func (ds GithubRESTAPI) getOrgRepos(ctx context.Context, token string, orgChan chan entity.Organization,
	reposBytesChan chan []byte, errorsChan chan error) {
	org := <-orgChan
	req, err := http.NewRequest("GET", org.ReposURL, nil)

	if err != nil {
		reposBytesChan <- nil
		errorsChan <- err
		return
	}

	req.Header.Set(OAUTH2_HEADER, OAUTH2_PREFIX+token)

	resp, err := ds.client.Do(req)

	if err != nil {
		reposBytesChan <- nil
		errorsChan <- err
		return
	}

	reposBody, err := bodyutil.ReadResponseBody(resp)

	if err != nil {
		reposBytesChan <- nil
		errorsChan <- err
		return
	}

	reposBytesChan <- reposBody
	errorsChan <- nil
}

// GetReposPullRequests returns body with orgs repos pulls.
func (ds GithubRESTAPI) GetReposPullRequests(ctx context.Context, token string, repos []entity.Repository) ([][]byte, error) {
	resultSetBytes := [][]byte{}

	repoChan := make(chan entity.Repository, len(repos))
	pullsBytesChan := make(chan []byte, 1)
	errorsChan := make(chan error, 1)

	var wg sync.WaitGroup
	wg.Add(len(repos))

	for _, repo := range repos {
		repoChan <- repo

		go ds.getReposPulls(ctx, token, repoChan, pullsBytesChan, errorsChan, &wg)

		pullsBody := <-pullsBytesChan

		err := <-errorsChan

		if err != nil {
			return nil, err
		}

		resultSetBytes = append(resultSetBytes, pullsBody)
	}

	wg.Wait()

	return resultSetBytes, nil
}

func (ds GithubRESTAPI) getReposPulls(ctx context.Context, token string, repoChan chan entity.Repository,
	pullsBytesChan chan []byte, errorsChan chan error, wg *sync.WaitGroup) {
	defer wg.Done()

	repo := <-repoChan

	req, err := http.NewRequest("GET", trimPullsURL(repo.PullsURL), nil)

	if err != nil {
		pullsBytesChan <- nil
		errorsChan <- err
		return
	}

	req.Header.Set(OAUTH2_HEADER, OAUTH2_PREFIX+token)

	resp, err := ds.client.Do(req)

	if err != nil {
		pullsBytesChan <- nil
		errorsChan <- err
	}

	pullsBody, err := bodyutil.ReadResponseBody(resp)

	if err != nil {
		pullsBytesChan <- nil
		errorsChan <- err
	}

	pullsBytesChan <- pullsBody
	errorsChan <- nil
}

func trimPullsURL(pullsUrls string) string {
	return strings.Trim(pullsUrls, "{/number/}")
}
