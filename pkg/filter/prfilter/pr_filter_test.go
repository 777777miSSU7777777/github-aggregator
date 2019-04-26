package prfilter

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/777777miSSU7777777/github-aggregator/pkg/entity"
)

var allPulls []entity.PullRequest
var testAssignee entity.Assignee
var testReviewer entity.Reviewer

func init() {
	for i := 0; i < 10; i++ {
		allPulls = append(allPulls, entity.PullRequest{Title: string(i)})
	}

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			allPulls[i].State = "open"
		}
	}

	testAssignee = entity.Assignee{Login: "test-user"}

	testReviewer = entity.Reviewer{Login: "test-user"}

	allPulls[0].Assignees = append(allPulls[0].Assignees, testAssignee)
	allPulls[9].Assignees = append(allPulls[9].Assignees, testAssignee)

	allPulls[4].RequestedReviewers = append(allPulls[4].RequestedReviewers, testReviewer)
}

func TestFilterByChoice__PresentAssignee__Equal(t *testing.T) {
	testPulls := []entity.PullRequest{}

	testPulls = append(testPulls, allPulls[0])
	testPulls = append(testPulls, allPulls[9])

	filterPulls := FilterByAssignee(allPulls, testAssignee)

	assert.Equal(t, testPulls, filterPulls)
}

func TestFilterByChoice__MissingAssignee__Equal(t *testing.T) {
	testPulls := []entity.PullRequest{}

	filterPulls := FilterByAssignee(allPulls, entity.Assignee{Login: "test-user-2"})

	assert.Equal(t, testPulls, filterPulls)
}

func TestFilterByChoice__PresentReviewer__Equal(t *testing.T) {
	testPulls := []entity.PullRequest{}

	testPulls = append(testPulls, allPulls[4])

	filterPulls := FilterByReviewer(allPulls, testReviewer)

	assert.Equal(t, testPulls, filterPulls)
}

func TestFilterByChoice__MissingReviewer__Equal(t *testing.T) {
	testPulls := []entity.PullRequest{}

	filterPulls := FilterByReviewer(allPulls, entity.Reviewer{Login: "test-user-2"})

	assert.Equal(t, testPulls, filterPulls)
}

func TestFilterByChoice__PresentParticipation__Equal(t *testing.T) {
	testPulls := []entity.PullRequest{}

	testPulls = append(testPulls, allPulls[0])
	testPulls = append(testPulls, allPulls[4])
	testPulls = append(testPulls, allPulls[9])

	filterPulls := FilterByParticipation(allPulls, entity.User{Login: "test-user"})

	assert.Equal(t, testPulls, filterPulls)
}

func TestFilterByChoice__MissingParticipation__Equal(t *testing.T) {
	testPulls := []entity.PullRequest{}

	filterPulls := FilterByParticipation(allPulls, entity.User{Login: "test-user-2"})

	assert.Equal(t, testPulls, filterPulls)
}

func TestFilterByChoice__Correct__Equal(t *testing.T) {
	testPulls := []entity.PullRequest{}

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			testPulls = append(testPulls, allPulls[i])
		}
	}

	filterPulls := FilterByState(allPulls, "open")

	assert.Equal(t, testPulls, filterPulls)
}

func TestFilterByChoice__NotCorrect__Equal(t *testing.T) {
	testPulls := []entity.PullRequest{}

	for i := 0; i < 10; i++ {
		if i%2 != 0 {
			testPulls = append(testPulls, allPulls[i])
		}
	}

	filterPulls := FilterByState(allPulls, "open")

	assert.NotEqual(t, testPulls, filterPulls)
}
