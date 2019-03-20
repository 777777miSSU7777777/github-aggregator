package userfactory

import (
	"encoding/json"

	"github.com/777777miSSU7777777/github-aggregator/pkg/entity"
)

func New(userBytes []byte)(*entity.User, error){
	user := entity.User{}

	err := json.Unmarshal(userBytes, &user)

	if err != nil {
		return nil, err
	}

	return &user, nil
} 