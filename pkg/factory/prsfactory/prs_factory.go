// Package prsfactory contains a factory for array of Pull Requests.
package prsfactory

import (
	"encoding/json"

	"github.com/777777miSSU7777777/github-aggregator/pkg/entity"
)

// New returns and array of Pull Requests.
// Byte array param "orgsBytes" responsible for pull requests data from pull requests query.
// If json.Unmarshal occurs any error, this will be returned.
func New(prsBytes []byte) ([]entity.PullRequest, error) {
	prs := []entity.PullRequest{}

	err := json.Unmarshal(prsBytes, &prs)

	if err != nil {
		return nil, err
	}

	return prs, nil
}
