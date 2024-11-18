package cli

import (
	"errors"
	"fmt"

	"github.com/DavidLSaldana/gator/internal/state"
)

type command struct {
	name string
	args []string
}

func handlerLogin(s *state.State, cmd command) error {
	if (len(cmd.args) == 0) || (len(cmd.args) > 1) {
		return errors.New("Command: Login expects a single argument username")
	}

	s.Cfg.CurrentUserName = cmd.args[0]

	fmt.Printf("user has been set to %s", s.Cfg.CurrentUserName)

	return nil
}

type commands struct {
	commandMap map[string]func(*state.State, command) error
}

func (c *commands) register(name string, f func(*state.State, command) error) {
	c.commandMap[name] = f
}

func (c *commands) run(s *state.State, cmd command) error {
	command, ok := c.commandMap[cmd.name]
	if !ok {
		return errors.New("Please enter a valid command")
	}
	//not sure about this, check it next run:
	command(s, cmd)
	return nil
}
