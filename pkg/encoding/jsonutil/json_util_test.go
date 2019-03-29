package jsonutil

import (
	"encoding/json"
	"testing"

	"github.com/777777miSSU7777777/github-aggregator/pkg/crypto/randutil"
	"github.com/stretchr/testify/assert"
)

var testMap map[string]interface{}

func tearUp() {
	testMap = map[string]interface{}{}
	testMap["1"] = "one"
	testMap["2"] = "two"
	testMap["3"] = "three"
}

func tearDown() {
	delete(testMap, "1")
	delete(testMap, "2")
	delete(testMap, "3")
	testMap = nil
}

func TestBytesToMap__SameMap__Successful(t *testing.T) {
	tearUp()

	jsonBytes, _ := json.Marshal(testMap)

	m, _ := BytesToMap(jsonBytes)

	assert.Equal(t, testMap, m)
}

func TestBytesToMap__DifferentMaps__Failed(t *testing.T) {
	tearUp()

	m := map[string]interface{}{}
	m["one"] = "1"
	m["two"] = "2"
	m["three"] = "3"

	jsonBytes, _ := json.Marshal(m)

	m, _ = BytesToMap(jsonBytes)

	assert.NotEqual(t, testMap, m)
}

func TestBytesToMap__RandomBytes__Error(t *testing.T) {
	randomBytes, _ := randutil.GenerateRandomBytes(16)

	_, err := BytesToMap(randomBytes)

	assert.Error(t, err)
}
