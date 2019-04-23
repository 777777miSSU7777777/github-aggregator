package session

import (
	"context"

	"github.com/777777miSSU7777777/github-aggregator/pkg/factory/orgsfactory"
	"github.com/777777miSSU7777777/github-aggregator/pkg/factory/userfactory"
	"github.com/777777miSSU7777777/github-aggregator/pkg/log"
	"github.com/777777miSSU7777777/github-aggregator/pkg/query"
)

var sessionService *SessionService

func init() {
	sessionService = &SessionService{}
}

// GetSessionService returns current session service.
func GetSessionService() *SessionService {
	return sessionService
}

// SessionService service which caches most static data.
type SessionService struct {
	currentSession *Session
}

// StartSession starts new session.
func (s *SessionService) StartSession(token string) {
	s.currentSession = &Session{}

	userBytes, err := query.GetDataSource().GetUser(context.Background(), token)

	if err != nil {
		log.Error.Println(err)
	}

	user, err := userfactory.New(userBytes)

	if err != nil {
		log.Error.Println(err)
	}

	s.currentSession.SetCurrentUser(*user)

	orgsBytes, err := query.GetDataSource().GetOrgs(context.Background(), token)

	if err != nil {
		log.Error.Println(err)
	}

	orgs, err := orgsfactory.New(orgsBytes)

	if err != nil {
		log.Error.Println(err)
	}

	s.currentSession.SetUserOrgs(*orgs)
}

// GetSession return a copy of current session.
func (s SessionService) GetSession() Session {
	return *s.currentSession
}

// TODO
func (s *SessionService) UpdateSession(session Session) {

}

// CloseSession closes current session.
func (s *SessionService) CloseSession(token string) {
	s.currentSession = nil
}
