package rest

import (
	"github.com/777777miSSU7777777/github-aggregator/pkg/entity"

	log "github.com/sirupsen/logrus"
)

type recoverMiddleware struct {
	logger *log.Logger
	next   RESTService
}

func WrapRecoverMiddleware(svc RESTService, logger *log.Logger) recoverMiddleware {
	return recoverMiddleware{logger, svc}
}

func (mw recoverMiddleware) CurrentUser() (user entity.User) {
	defer func() {
		err := recover()
		if err != nil {
			mw.logger.Warnf("Recovered from %s", err)
		}
	}()

	user = mw.next.CurrentUser()
	return
}

func (mw recoverMiddleware) TokenScopes() (scopes []entity.Scope, err error) {
	defer func() {
		err := recover()
		if err != nil {
			mw.logger.Warnf("Recovered from %s", err)
		}
	}()

	scopes, err = mw.next.TokenScopes()
	return
}

func (mw recoverMiddleware) UserOrgs() (orgs []entity.Organization) {
	defer func() {
		err := recover()
		if err != nil {
			mw.logger.Warnf("Recovered from %s", err)
		}
	}()

	orgs = mw.next.UserOrgs()
	return
}

func (mw recoverMiddleware) FilteredPulls(filter string, orgsChoice []string) (pulls []entity.PullRequest, err error) {
	defer func() {
		err := recover()
		if err != nil {
			mw.logger.Warnf("Recovered from %s", err)
		}
	}()

	pulls, err = mw.next.FilteredPulls(filter, orgsChoice)
	return
}
