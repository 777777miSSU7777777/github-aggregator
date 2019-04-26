// Package prsfactory contains a factory for array of Pull Requests.
package prsfactory

import (
	"encoding/json"

	"github.com/777777miSSU7777777/github-aggregator/pkg/entity"
)

// New returns an array of Pull Requests.
func New(prsBytes [][]byte) ([]entity.PullRequest, error) {
	prs := []entity.PullRequest{}

	for _, repoPRsBytes := range prsBytes {
		repoPRs := []entity.PullRequest{}
		err := json.Unmarshal(repoPRsBytes, &repoPRs)

		if err != nil {
			return nil, err
		}

		prs = append(prs, repoPRs...)
	}

	return prs, nil
}
