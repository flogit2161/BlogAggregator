package main

import (
	"context"
	"errors"

	"github.com/flogit2161/BlogAggregator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return errors.New("No user logged in at the moment")
		}
		return handler(s, cmd, user)
	}
}
