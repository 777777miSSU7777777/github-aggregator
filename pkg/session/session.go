package session

import (
	"github.com/777777miSSU7777777/github-aggregator/pkg/entity"
)

type Session struct {
	currentUser  entity.User
	userOrgs     []entity.Organization
	selectedOrgs []string
}

func (s Session) GetCurrentUser() entity.User {
	return s.currentUser
}

func (s *Session) SetCurrentUser(newUser entity.User) {
	s.currentUser = newUser
}

func (s Session) GetUserOrgs() []entity.Organization {
	return s.userOrgs
}

func (s *Session) SetUserOrgs(newOrgs []entity.Organization) {
	s.userOrgs = newOrgs
}
