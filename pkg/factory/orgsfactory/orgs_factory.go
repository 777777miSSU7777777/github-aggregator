// Package orgsfactory contains a factory for array of Organizations.
package orgsfactory

import (
	"encoding/json"

	"github.com/777777miSSU7777777/github-aggregator/pkg/entity"
)

// New returns an array of Organizations.
func New(orgsBytes []byte) ([]entity.Organization, error) {
	orgs := []entity.Organization{}

	err := json.Unmarshal(orgsBytes, &orgs)

	if err != nil {
		return nil, err
	}

	return orgs, nil
}
