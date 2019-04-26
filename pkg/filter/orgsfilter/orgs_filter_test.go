package orgsfilter

import (
	"testing"

	"github.com/777777miSSU7777777/github-aggregator/pkg/entity"
	"github.com/stretchr/testify/assert"
)

var allOrgs []entity.Organization

func init() {
	for i := 0; i < 10; i++ {
		allOrgs = append(allOrgs, entity.Organization{Login: string(i)})
	}
}

func TestFilterByChoice__OrgsChoiceNotEmpty__Equal(t *testing.T) {
	orgsChoice := []string{}

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			orgsChoice = append(orgsChoice, string(i))
		}
	}

	testOrgs := []entity.Organization{}

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			testOrgs = append(testOrgs, entity.Organization{Login: string(i)})
		}
	}

	filterOrgs := FilterByChoice(allOrgs, orgsChoice)

	assert.Equal(t, testOrgs, filterOrgs)
}

func TestFilterByChoice__OrgsChoiceEmpty__Equal(t *testing.T) {
	orgsChoice := []string{}

	testOrgs := []entity.Organization{}

	filterOrgs := FilterByChoice(allOrgs, orgsChoice)

	assert.Equal(t, testOrgs, filterOrgs)
}
