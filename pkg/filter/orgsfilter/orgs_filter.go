package orgsfilter

import (
	"github.com/777777miSSU7777777/github-aggregator/pkg/entity"
)

func FilterByChoice(orgs []entity.Organization, orgsChoice []string) []entity.Organization {
	resultSet := []entity.Organization{}

	for _, org := range orgs {
		if orgIndex(orgsChoice, org) != -1 {
			resultSet = append(resultSet, org)
		}
	}

	return resultSet
}

func orgIndex(orgsChoice []string, org entity.Organization) int {
	for i, selectedOrg := range orgsChoice {
		if org.Login == selectedOrg {
			return i
		}
	}

	return -1
}
