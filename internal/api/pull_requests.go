package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/777777miSSU7777777/github-aggregator/internal/security/tokenservice"
	"github.com/777777miSSU7777777/github-aggregator/pkg/entity"

	"github.com/777777miSSU7777777/github-aggregator/pkg/factory/prsfactory"
	"github.com/777777miSSU7777777/github-aggregator/pkg/factory/reposfactory"
	"github.com/777777miSSU7777777/github-aggregator/pkg/filter/orgsfilter"
	"github.com/777777miSSU7777777/github-aggregator/pkg/filter/prfilter"
	"github.com/777777miSSU7777777/github-aggregator/pkg/log"
	"github.com/777777miSSU7777777/github-aggregator/pkg/query"
	"github.com/777777miSSU7777777/github-aggregator/pkg/session"
)

const (
	// FILTER_PARAM pulls filter param.
	FILTER_PARAM = "filter"

	// ORGS_CHOICE_PARAM param for chosen orgs.
	ORGS_CHOICE_PARAM = "orgs_choice"

	// ALL filter type for pulls assigned or review requested to user.
	ALL = "all"

	// ASSIGNEE filter type for pulls assigned to user.
	ASSIGNEE = "assignee"

	// REVIEWER type for pulls review requested to user.
	REVIEWER = "reviewer"

	// OPEN type for open pulls.
	OPEN = "open"
)

// PullRequests returns response with pull requests for chosen filter.
func PullRequests(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	token := tokenservice.GetToken()

	orgs := session.GetSessionService().GetSession().GetUserOrgs()

	orgsChoiceBody := req.FormValue(ORGS_CHOICE_PARAM)

	orgsChoice := []string{}

	err := json.Unmarshal([]byte(orgsChoiceBody), &orgsChoice)

	if err != nil {
		log.Error.Println(err)
		http.Error(rw, "Internal server error", http.StatusInternalServerError)
	}

	orgs = orgsfilter.FilterByChoice(orgs, orgsChoice)

	reposBytes, err := query.GetDataSource().GetOrgsRepos(context.Background(), token, orgs)

	repos, err := reposfactory.New(reposBytes)

	if err != nil {
		log.Error.Println(err)
		http.Error(rw, "Internal server error", http.StatusInternalServerError)
	}

	pullRequestsBytes, err := query.GetDataSource().GetReposPullRequests(context.Background(), token, repos)

	if err != nil {
		log.Error.Println(err)
		http.Error(rw, "Internal server error", http.StatusInternalServerError)
	}

	pullRequests, err := prsfactory.New(pullRequestsBytes)

	pullRequests = prfilter.FilterByState(pullRequests, OPEN)

	if err != nil {
		log.Error.Println(err)
		http.Error(rw, "Internal server error", http.StatusInternalServerError)
	}

	user := session.GetSessionService().GetSession().GetCurrentUser()

	switch filter := req.FormValue(FILTER_PARAM); filter {
	case ALL:
		pullRequests = prfilter.FilterByParticipation(pullRequests, user)
	case ASSIGNEE:
		pullRequests = prfilter.FilterByAssignee(pullRequests, entity.Assignee(user))
	case REVIEWER:
		pullRequests = prfilter.FilterByReviewer(pullRequests, entity.Reviewer(user))
	}

	err = json.NewEncoder(rw).Encode(pullRequests)

	if err != nil {
		log.Error.Println(err)
		http.Error(rw, "Internal server error", http.StatusInternalServerError)
	} else {
		log.Info.Printf("Filtered by %s and sent %s", req.FormValue(FILTER_PARAM), pullRequests)
	}
}
