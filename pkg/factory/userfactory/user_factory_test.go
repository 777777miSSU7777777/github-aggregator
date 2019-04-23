package userfactory

import (
	"encoding/json"
	"testing"

	"github.com/777777miSSU7777777/github-aggregator/pkg/crypto/randutil"
	"github.com/777777miSSU7777777/github-aggregator/pkg/entity"
	"github.com/stretchr/testify/assert"
)

func TestNew__CorrectBytes__Equal(t *testing.T) {
	testUser := entity.User{}

	testUser.Login = "test_user"

	jsonBytes, _ := json.Marshal(testUser)

	factoryUser, _ := New(jsonBytes)

	assert.ObjectsAreEqual(testUser, factoryUser)
}

func TestNew__IncorrectBytes__Equal(t *testing.T) {
	randomBytes, _ := randutil.GenerateRandomBytes(16)

	_, err := New(randomBytes)

	assert.Error(t, err)
}
