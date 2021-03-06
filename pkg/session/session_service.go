package session

import (
	"context"

	"github.com/777777miSSU7777777/github-aggregator/pkg/factory/orgsfactory"
	"github.com/777777miSSU7777777/github-aggregator/pkg/factory/userfactory"
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
func (s *SessionService) StartSession(token string) error {
	s.currentSession = &Session{}

	userBytes, err := query.GetDataSource().GetUser(context.Background(), token)

	if err != nil {
		return err
	}

	user, err := userfactory.New(userBytes)

	if err != nil {
		return err
	}

	s.currentSession.SetCurrentUser(*user)

	orgsBytes, err := query.GetDataSource().GetOrgs(context.Background(), token)

	if err != nil {
		return err
	}

	orgs, err := orgsfactory.New(orgsBytes)

	if err != nil {
		return err
	}

	s.currentSession.SetUserOrgs(orgs)

	return nil
}

// GetSession return a copy of current session.
func (s SessionService) GetSession() Session {
	return *s.currentSession
}

// HasActiveSession returns state of session state (active/inactive).
func (s SessionService) HasActiveSession() bool {
	return s.currentSession != nil
}

// UpdateSession TODO.
func (s *SessionService) UpdateSession(session Session) {

}

// CloseSession closes current session.
func (s *SessionService) CloseSession() {
	s.currentSession = nil
}
