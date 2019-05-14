package rest

import (
	"github.com/777777miSSU7777777/github-aggregator/pkg/entity"

	log "github.com/sirupsen/logrus"
)

type loggingMiddleware struct {
	logger *log.Logger
	next   RESTService
}

func WrapLoggingMiddleware(svc RESTService, logger *log.Logger) loggingMiddleware {
	return loggingMiddleware{logger, svc}
}

func (mw loggingMiddleware) CurrentUser() (user entity.User) {
	defer func() {
		mw.logger.WithFields(log.Fields{
			"user":  user,
			"error": nil,
		}).Infoln()
	}()

	user = mw.next.CurrentUser()
	return
}

func (mw loggingMiddleware) TokenScopes() (scopes []entity.Scope, err error) {
	defer func() {
		mw.logger.WithFields(log.Fields{
			"scopes": scopes,
			"error":  err,
		}).Infoln()
	}()

	scopes, err = mw.next.TokenScopes()
	return
}

func (mw loggingMiddleware) UserOrgs() (orgs []entity.Organization) {
	defer func() {
		mw.logger.WithFields(log.Fields{
			"orgs":  orgs,
			"error": nil,
		}).Infoln()
	}()

	orgs = mw.next.UserOrgs()
	return
}

func (mw loggingMiddleware) FilteredPulls(filter string, orgsChoice []string) (pulls []entity.PullRequest, err error) {
	defer func() {
		mw.logger.WithFields(log.Fields{
			"filter":      filter,
			"orgs_choice": orgsChoice,
			"pulls":       pulls,
			"error":       err,
		}).Infoln()
	}()

	pulls, err = mw.next.FilteredPulls(filter, orgsChoice)
	return
}
