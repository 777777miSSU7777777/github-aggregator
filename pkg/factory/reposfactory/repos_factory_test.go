package reposfactory

import (
	"encoding/json"
	"testing"

	"github.com/777777miSSU7777777/github-aggregator/pkg/crypto/randutil"
	"github.com/777777miSSU7777777/github-aggregator/pkg/entity"
	"github.com/stretchr/testify/assert"
)

func TestNew__CorrectBytes__Equal(t *testing.T) {
	testRepos := []entity.Repository{}

	testRepo1 := entity.Repository{Name: "test-repo-1"}

	testRepo2 := entity.Repository{Name: "test-repo-2"}

	testRepos = append(testRepos, testRepo1)

	testRepos = append(testRepos, testRepo2)

	reposJsonBytes, _ := json.Marshal(testRepos)

	jsonBytes := [][]byte{}

	jsonBytes = append(jsonBytes, reposJsonBytes)

	factoryRepos, _ := New(jsonBytes)

	assert.ObjectsAreEqual(testRepos, factoryRepos)
}

func TestNew__IncorrectBytes__Error(t *testing.T) {
	randomBytes, _ := randutil.GenerateRandomBytes(16)

	randomBytesArray := [][]byte{}

	randomBytesArray = append(randomBytesArray, randomBytes)

	_, err := New(randomBytesArray)

	assert.Error(t, err)
}
