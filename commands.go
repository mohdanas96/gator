package main

import (
	"errors"
	"fmt"

	"github.com/mohdanas96/gator/internal/config"
)

type state struct {
	c *config.Config
}

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
	s.c.SetUser(name)

	fmt.Printf("%v is now using gator\n", name)

	return nil
}
