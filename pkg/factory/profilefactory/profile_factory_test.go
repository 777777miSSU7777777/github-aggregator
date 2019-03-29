package profilefactory

import (
	"encoding/json"
	"testing"

	"github.com/777777miSSU7777777/github-aggregator/pkg/crypto/randutil"
	"github.com/777777miSSU7777777/github-aggregator/pkg/entity"
	"github.com/stretchr/testify/assert"
)

func TestNew__CorrectBytes__Equal(t *testing.T) {
	testProfile := entity.Profile{}

	testProfile.Username = "test user"
	testProfile.ProfileURL = "test.com/user"
	testProfile.AvatarURL = "test/user/profile_pic"

	testScopes := []string{"user", "repository"}

	testProfile.Scopes = testScopes

	jsonBytes, _ := json.Marshal(testProfile)

	factoryProfile, _ := New(jsonBytes, testScopes)

	assert.ObjectsAreEqual(testProfile, factoryProfile)
}

func TestNew__IncorrectBytes__Error(t *testing.T) {
	randomBytes, _ := randutil.GenerateRandomBytes(16)

	testScopes := []string{"user", "repository"}

	_, err := New(randomBytes, testScopes)

	assert.Error(t, err)
}
