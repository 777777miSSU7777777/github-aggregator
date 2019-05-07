package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/777777miSSU7777777/github-aggregator/pkg/entity"
	"github.com/777777miSSU7777777/github-aggregator/pkg/filter/prfilter"

	"github.com/stretchr/testify/assert"
)

func TestPullRequests_FilterByAll(t *testing.T) {
	original := PullRequests

	testUser := entity.User{Login: "test-user"}

	testPull1 := entity.PullRequest{Assignees: []entity.Assignee{entity.Assignee(testUser)}}

	testPull2 := entity.PullRequest{RequestedReviewers: []entity.Reviewer{entity.Reviewer(testUser)}}

	testPulls := []entity.PullRequest{testPull1, testPull2}

	testPullsJSONBytes, _ := json.Marshal(testPulls)
	testPullsJSONBytes = append(testPullsJSONBytes, 0xa)

	PullRequests = func(rw http.ResponseWriter, req *http.Request) {
		switch filter := req.FormValue(FILTER_PARAM); filter {
		case ALL:
			testPulls = prfilter.FilterByParticipation(testPulls, testUser)
		case ASSIGNEE:
			testPulls = prfilter.FilterByAssignee(testPulls, entity.Assignee(testUser))
		case REVIEWER:
			testPulls = prfilter.FilterByReviewer(testPulls, entity.Reviewer(testUser))
		}

		json.NewEncoder(rw).Encode(testPulls)
	}

	defer func() {
		PullRequests = original
	}()

	req, _ := http.NewRequest("GET", "/api/pulls?filter=all&orgs_choice=['test-org']", nil)

	rw := httptest.NewRecorder()

	handler := http.HandlerFunc(PullRequests)

	handler.ServeHTTP(rw, req)

	responseBody := rw.Body.Bytes()

	assert.Equal(t, testPullsJSONBytes, responseBody)
}

func TestPullRequests_FilterByAssignee(t *testing.T) {
	original := PullRequests

	testUser := entity.User{Login: "test-user"}

	testPull1 := entity.PullRequest{Assignees: []entity.Assignee{entity.Assignee(testUser)}}

	testPull2 := entity.PullRequest{RequestedReviewers: []entity.Reviewer{entity.Reviewer(testUser)}}

	testPulls := []entity.PullRequest{testPull1, testPull2}

	testPullsAssignee := []entity.PullRequest{testPull1}

	testPullsJSONBytes, _ := json.Marshal(testPullsAssignee)
	testPullsJSONBytes = append(testPullsJSONBytes, 0xa)

	PullRequests = func(rw http.ResponseWriter, req *http.Request) {
		switch filter := req.FormValue(FILTER_PARAM); filter {
		case ALL:
			testPulls = prfilter.FilterByParticipation(testPulls, testUser)
		case ASSIGNEE:
			testPulls = prfilter.FilterByAssignee(testPulls, entity.Assignee(testUser))
		case REVIEWER:
			testPulls = prfilter.FilterByReviewer(testPulls, entity.Reviewer(testUser))
		}

		json.NewEncoder(rw).Encode(testPulls)
	}

	defer func() {
		PullRequests = original
	}()

	req, _ := http.NewRequest("GET", "/api/pulls?filter=assignee&orgs_choice=['test-org']", nil)

	rw := httptest.NewRecorder()

	handler := http.HandlerFunc(PullRequests)

	handler.ServeHTTP(rw, req)

	responseBody := rw.Body.Bytes()

	assert.Equal(t, testPullsJSONBytes, responseBody)
}

func TestPullRequests_FilterByReviewer(t *testing.T) {
	original := PullRequests

	testUser := entity.User{Login: "test-user"}

	testPull1 := entity.PullRequest{Assignees: []entity.Assignee{entity.Assignee(testUser)}}

	testPull2 := entity.PullRequest{RequestedReviewers: []entity.Reviewer{entity.Reviewer(testUser)}}

	testPulls := []entity.PullRequest{testPull1, testPull2}

	testPullsReviewer := []entity.PullRequest{testPull2}

	testPullsJSONBytes, _ := json.Marshal(testPullsReviewer)
	testPullsJSONBytes = append(testPullsJSONBytes, 0xa)

	PullRequests = func(rw http.ResponseWriter, req *http.Request) {
		switch filter := req.FormValue(FILTER_PARAM); filter {
		case ALL:
			testPulls = prfilter.FilterByParticipation(testPulls, testUser)
		case ASSIGNEE:
			testPulls = prfilter.FilterByAssignee(testPulls, entity.Assignee(testUser))
		case REVIEWER:
			testPulls = prfilter.FilterByReviewer(testPulls, entity.Reviewer(testUser))
		}

		json.NewEncoder(rw).Encode(testPulls)
	}

	defer func() {
		PullRequests = original
	}()

	req, _ := http.NewRequest("GET", "/api/pulls?filter=reviewer&orgs_choice=['test-org']", nil)

	rw := httptest.NewRecorder()

	handler := http.HandlerFunc(PullRequests)

	handler.ServeHTTP(rw, req)

	responseBody := rw.Body.Bytes()

	assert.Equal(t, testPullsJSONBytes, responseBody)
}
