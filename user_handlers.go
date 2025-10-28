package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mohdanas96/gator/internal/database"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return errors.New("username is required")
	}

	name := cmd.args[0]

	user, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		return errors.New("user doesn't exist")
	}

	s.c.SetUser(user.Name)

	fmt.Printf("%v is now using gator\n", name)

	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return errors.New("username is required")
	}
	name := cmd.args[0]

	args := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
	}

	user, err := s.db.CreateUser(context.Background(), args)
	if err != nil {
		return err
	}

	s.c.SetUser(user.Name)

	return nil
}

func handlerDeleteUsers(s *state, _ command) error {
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		return err
	}

	s.c.SetUser("")

	fmt.Println("Deleted users successfully")

	return nil
}

func handlerGetAllUsers(s *state, _ command) error {
	users, err := s.db.GetAllUsers(context.Background())
	if err != nil {
		return err
	}

	if len(users) == 0 {
		fmt.Println("Currently no users. Register as a user using 'register' command")
		return nil
	}

	for _, user := range users {
		if user.Name == s.c.Current_user_name {
			fmt.Printf("* %s (current)\n", user.Name)
		} else {
			fmt.Printf("* %s\n", user.Name)
		}
	}

	return nil
}
