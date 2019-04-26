package prsfactory

import (
	"encoding/json"
	"testing"

	"github.com/777777miSSU7777777/github-aggregator/pkg/crypto/randutil"
	"github.com/777777miSSU7777777/github-aggregator/pkg/entity"
	"github.com/stretchr/testify/assert"
)

func TestNew__CorrectBytes__Equal(t *testing.T) {
	testPulls := []entity.PullRequest{}

	testPull1 := entity.PullRequest{Title: "test-pr-1"}

	testPull2 := entity.PullRequest{Title: "test-pr-2"}

	testPulls = append(testPulls, testPull1)

	testPulls = append(testPulls, testPull2)

	pullsJsonBytes, _ := json.Marshal(testPulls)

	jsonBytes := [][]byte{}

	jsonBytes = append(jsonBytes, pullsJsonBytes)

	factoryPulls, _ := New(jsonBytes)

	assert.ObjectsAreEqual(testPulls, factoryPulls)
}

func TestNew__IncorrectBytes__Error(t *testing.T) {
	randomBytes, _ := randutil.GenerateRandomBytes(16)

	randomBytesArray := [][]byte{}

	randomBytesArray = append(randomBytesArray, randomBytes)

	_, err := New(randomBytesArray)

	assert.Error(t, err)
}
