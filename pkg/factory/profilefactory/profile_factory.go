// Package profilefactory contains a factory for profile entity.
package profilefactory

import (
	"encoding/json"

	"github.com/777777miSSU7777777/github-aggregator/pkg/entity"
)

// New returns an instance of Profile.
// Byte array param "userBytes" is responsible for user data from user query.
// String array param "scopes" is responsible for Github API scopes for provided token from scopes query.
// If json.Unmarshal occurs any error, this will be returned.
func New(userBytes []byte, scopes []string) (*entity.Profile, error) {
	profile := entity.Profile{}

	err := json.Unmarshal(userBytes, &profile)

	if err != nil {
		return nil, err
	}

	profile.Scopes = scopes

	return &profile, nil
}
