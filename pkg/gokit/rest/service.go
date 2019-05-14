package rest

import (
	"context"

	"github.com/777777miSSU7777777/github-aggregator/pkg/entity"
	"github.com/777777miSSU7777777/github-aggregator/pkg/factory/prsfactory"
	"github.com/777777miSSU7777777/github-aggregator/pkg/factory/reposfactory"
	"github.com/777777miSSU7777777/github-aggregator/pkg/filter/orgsfilter"
	"github.com/777777miSSU7777777/github-aggregator/pkg/filter/prfilter"
	"github.com/777777miSSU7777777/github-aggregator/pkg/query/datasource"
	"github.com/777777miSSU7777777/github-aggregator/pkg/session"
	"github.com/777777miSSU7777777/github-aggregator/pkg/token"
)

const (
	// ALL filter type for pulls assigned or review requested to user.
	ALL = "all"

	// ASSIGNEE filter type for pulls assigned to user.
	ASSIGNEE = "assignee"

	// REVIEWER type for pulls review requested to user.
	REVIEWER = "reviewer"
)

type RESTService interface {
	CurrentUser() entity.User
	TokenScopes() ([]entity.Scope, error)
	UserOrgs() []entity.Organization
	FilteredPulls(string, []string) ([]entity.PullRequest, error)
}

type service struct {
	repository datasource.DataSource
}

func NewRestServiceImpl() RESTService {
	return &service{
		repository: datasource.NewGithubRESTAPI(),
	}
}

func (s service) CurrentUser() entity.User {
	return session.GetSessionService().GetSession().GetCurrentUser()
}

func (s service) TokenScopes() ([]entity.Scope, error) {
	tkn := token.GetTokenService().GetToken()
	scopes, err := s.repository.GetScopes(context.Background(), tkn)

	if err != nil {
		return nil, err
	}

	return scopes, nil
}

func (s service) UserOrgs() []entity.Organization {
	return session.GetSessionService().GetSession().GetUserOrgs()
}

func (s service) FilteredPulls(filter string, orgsChoice []string) ([]entity.PullRequest, error) {
	orgs := session.GetSessionService().GetSession().GetUserOrgs()

	orgs = orgsfilter.FilterByChoice(orgs, orgsChoice)

	tkn := token.GetTokenService().GetToken()

	reposBytes, err := s.repository.GetOrgsRepos(context.Background(), tkn, orgs)

	if err != nil {
		return nil, err
	}

	repos, err := reposfactory.New(reposBytes)

	if err != nil {
		return nil, err
	}

	pullsBytes, err := s.repository.GetReposPullRequests(context.Background(), tkn, repos)

	if err != nil {
		return nil, err
	}

	pulls, err := prsfactory.New(pullsBytes)

	if err != nil {
		return nil, err
	}

	user := session.GetSessionService().GetSession().GetCurrentUser()

	switch filter {
	case ALL:
		pulls = prfilter.FilterByParticipation(pulls, user)
	case ASSIGNEE:
		pulls = prfilter.FilterByAssignee(pulls, entity.Assignee(user))
	case REVIEWER:
		pulls = prfilter.FilterByReviewer(pulls, entity.Reviewer(user))
	}

	return pulls, nil
}
