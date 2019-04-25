// Package reposfactory contains a factory for array of Repositories.
package reposfactory

import (
	"encoding/json"

	"github.com/777777miSSU7777777/github-aggregator/pkg/entity"
)

// New returns and array of Repositories.
// Byte array param "orgsBytes" responsible for repositores data from repositories query.
// If json.Unmarshal occurs any error, this will be returned.
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
