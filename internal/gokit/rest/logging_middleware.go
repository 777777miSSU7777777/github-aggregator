package rest

import (
	"time"

	"github.com/777777miSSU7777777/github-aggregator/pkg/entity"
	"github.com/777777miSSU7777777/github-aggregator/pkg/time/timeutil"

	"github.com/go-kit/kit/log"
)

type loggingMiddleware struct {
	logger log.Logger
	next   RESTService
}

func WrapLoggingMiddleware(svc RESTService, logger log.Logger) loggingMiddleware {
	return loggingMiddleware{logger, svc}
}

func (mw loggingMiddleware) CurrentUser() (user entity.User) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "CurrentUser",
			"time", timeutil.GetCurrentTime(),
			"user", user,
			"err", nil,
			"took", time.Since(begin),
		)
	}(time.Now())

	user = mw.next.CurrentUser()
	return
}

func (mw loggingMiddleware) TokenScopes() (scopes []entity.Scope, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "TokenScopes",
			"time", timeutil.GetCurrentTime(),
			"scopes", scopes,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	scopes, err = mw.next.TokenScopes()
	return
}

func (mw loggingMiddleware) UserOrgs() (orgs []entity.Organization) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "UserOrgs",
			"time", timeutil.GetCurrentTime(),
			"orgs", orgs,
			"err", nil,
			"took", time.Since(begin),
		)
	}(time.Now())

	orgs = mw.next.UserOrgs()
	return
}

func (mw loggingMiddleware) FilteredPulls(filter string, orgsChoice []string) (pulls []entity.PullRequest, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "FilteredPulls",
			"time", timeutil.GetCurrentTime(),
			"filter", filter,
			"orgs_choice", orgsChoice,
			"pulls", pulls,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	pulls, err = mw.next.FilteredPulls(filter, orgsChoice)
	return
}
