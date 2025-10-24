package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mohdanas96/gator/internal/database"
)

type command struct {
	name string
	args []string
}

type commandHandler func(*state, command) error

type commands struct {
	commandRegistry map[string]commandHandler
}

func (c *commands) run(s *state, cmd command) error {
	fn := c.commandRegistry[cmd.name]
	err := fn(s, cmd)
	if err != nil {
		return err
	}
	return nil
}

func (c *commands) register(name string, f func(*state, command) error) error {
	_, ok := c.commandRegistry[name]
	if !ok {
		c.commandRegistry[name] = f
		return nil
	}
	return nil
}

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
