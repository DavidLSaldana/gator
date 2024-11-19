package cli

import (
	"errors"
	"fmt"

	"github.com/DavidLSaldana/gator/internal/state"
)

type Command struct {
	name string
	args []string
}

func handlerLogin(s *state.State, cmd Command) error {
	if (len(cmd.args) == 0) || (len(cmd.args) > 1) {
		return errors.New("Command: Login expects a single argument username")
	}

	s.Cfg.CurrentUserName = cmd.args[0]

	fmt.Printf("user has been set to %s", s.Cfg.CurrentUserName)

	return nil
}

type Commands struct {
	CommandMap map[string]func(*state.State, Command) error
}

func (c *Commands) Register(name string, f func(*state.State, Command) error) {
	c.CommandMap[name] = f
}

func (c *Commands) run(s *state.State, cmd Command) error {
	command, ok := c.CommandMap[cmd.name]
	if !ok {
		return errors.New("Please enter a valid command")
	}
	//not sure about this, check it next run:
	command(s, cmd)
	return nil
}
