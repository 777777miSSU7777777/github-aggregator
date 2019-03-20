package profilefactory

import (
	"encoding/json"

	"github.com/777777miSSU7777777/github-aggregator/pkg/entity"
)

func New(userBytes []byte, scopes []string)(*entity.Profile, error){
	profile := entity.Profile{}

	err := json.Unmarshal(userBytes, &profile)

	if err != nil {
		return nil, err
	}

	profile.Scopes = scopes

	return &profile, nil
} 