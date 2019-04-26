// Package reposfactory contains a factory for array of Repositories.
package reposfactory

import (
	"encoding/json"

	"github.com/777777miSSU7777777/github-aggregator/pkg/entity"
)

// New returns new array of Repositories.
func New(reposBytes [][]byte) ([]entity.Repository, error) {
	repos := []entity.Repository{}

	for _, orgsReposBytes := range reposBytes {
		orgRepos := []entity.Repository{}

		err := json.Unmarshal(orgsReposBytes, &orgRepos)

		if err != nil {
			return nil, err
		}

		repos = append(repos, orgRepos...)
	}

	return repos, nil
}
