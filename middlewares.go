package main

import (
	"context"
	"fmt"

	"github.com/mohdanas96/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		user, err := s.db.GetUser(context.Background(), s.c.Current_user_name)
		if err != nil {
			return fmt.Errorf("something went wrong while fetching user :: %v", err)
		}

		return handler(s, cmd, user)
	}
}
