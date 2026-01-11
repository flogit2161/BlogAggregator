package main

import (
	"errors"
	"fmt"

	"github.com/flogit2161/BlogAggregator/internal/config"
)

type state struct {
	cfg *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	handlers map[string]func(*state, command) error
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return errors.New("Please only use 1 argument to use command function")
	}
	err := s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return errors.New("Error setting up the user's name")
	}

	fmt.Println("User has been set")
	return nil
}

func (c *commands) run(s *state, cmd command) error {
	handler, exists := c.handlers[cmd.name]
	if !exists {
		return errors.New("Command doesnt exist yet")
	}
	return handler(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.handlers[name] = f
}
