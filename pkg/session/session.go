package session

import (
	"github.com/777777miSSU7777777/github-aggregator/pkg/entity"
)

// Session is entity for current user session.
type Session struct {
	currentUser  entity.User
	userOrgs     []entity.Organization
	selectedOrgs []string
}

// GetCurrentUser returns current user.
func (s Session) GetCurrentUser() entity.User {
	return s.currentUser
}

// SetCurrentUser sets current user.
func (s *Session) SetCurrentUser(newUser entity.User) {
	s.currentUser = newUser
}

// GetUserOrgs returns orgs of current user.
func (s Session) GetUserOrgs() []entity.Organization {
	return s.userOrgs
}

// SetUserOrgs sets orgs of current user.
func (s *Session) SetUserOrgs(newOrgs []entity.Organization) {
	s.userOrgs = newOrgs
}
