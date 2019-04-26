package orgsfactory

import (
	"encoding/json"
	"testing"

	"github.com/777777miSSU7777777/github-aggregator/pkg/crypto/randutil"
	"github.com/777777miSSU7777777/github-aggregator/pkg/entity"
	"github.com/stretchr/testify/assert"
)

func TestNew__CorrectBytes__Equal(t *testing.T) {
	testOrganizations := []entity.Organization{}

	testOrg1 := entity.Organization{Login: "testOrg1"}

	testOrg2 := entity.Organization{Login: "testOrg2"}

	testOrganizations = append(testOrganizations, testOrg1)

	testOrganizations = append(testOrganizations, testOrg2)

	jsonBytes, _ := json.Marshal(testOrganizations)

	factoryOrganizations, _ := New(jsonBytes)

	assert.ObjectsAreEqual(testOrganizations, factoryOrganizations)
}

func TestNew__IncorrectBytes__Error(t *testing.T) {
	randomBytes, _ := randutil.GenerateRandomBytes(16)

	_, err := New(randomBytes)

	assert.Error(t, err)
}
