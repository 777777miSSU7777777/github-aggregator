package datasrcfactory

import (
	"testing"

	"github.com/777777miSSU7777777/github-aggregator/pkg/query/datasource"
	"github.com/stretchr/testify/assert"
)

func TestNew__rest_api__Equal(t *testing.T) {
	testDatasrc := datasource.GithubRESTAPI{}

	factoryDataSrc := New("rest-api")

	assert.ObjectsAreEqual(testDatasrc, factoryDataSrc)
}

func TestNew__default__Equal(t *testing.T) {
	testDatasrc := datasource.GithubRESTAPI{}

	factoryDataSrc := New("default")

	assert.ObjectsAreEqual(testDatasrc, factoryDataSrc)
}
