package jsonutil

import (
	"encoding/json"
	"testing"

	"github.com/777777miSSU7777777/github-aggregator/pkg/crypto/randutil"
	"github.com/777777miSSU7777777/github-aggregator/pkg/log"
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

	jsonBytes, err := json.Marshal(testMap)

	if err != nil {
		log.Error.Fatalln(err)
	}

	m, err := BytesToMap(jsonBytes)

	if err != nil {
		log.Error.Fatalln(err)
	}

	assert.Equal(t, testMap, m)
}

func TestBytesToMap__DifferentMaps__Failed(t *testing.T) {
	tearUp()

	m := map[string]interface{}{}
	m["one"] = "1"
	m["two"] = "2"
	m["three"] = "3"

	jsonBytes, err := json.Marshal(m)

	if err != nil {
		log.Error.Fatalln(err)
	}

	m, err = BytesToMap(jsonBytes)

	if err != nil {
		log.Error.Fatalln(err)
	}

	assert.NotEqual(t, testMap, m)
}

func TestBytesToMap__RandomBytes__Error(t *testing.T) {
	randomBytes, err := randutil.GenerateRandomBytes(16)

	if err != nil {
		log.Error.Fatalln(err)
	}

	_, err = BytesToMap(randomBytes)

	assert.Error(t, err)
}
