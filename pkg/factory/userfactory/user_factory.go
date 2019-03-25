package userfactory

import (
	"encoding/json"

	"github.com/777777miSSU7777777/github-aggregator/pkg/entity"
)

// New returns and instance of User.
// Byte array param "userBytes" responsible for user data from user query.
// If json.Unmarshal occurs any error, this will be returned.
func New(userBytes []byte) (*entity.User, error) {
	user := entity.User{}

	err := json.Unmarshal(userBytes, &user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
