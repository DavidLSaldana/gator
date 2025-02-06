package cli

import (
	"errors"
	"fmt"
	"github.com/DavidLSaldana/gator/internal/config"
)

type state struct {
	cfgPointer *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	cmds map[string]func(*state, command) error
}

// registers a new handler function for a command name
func (c *commands) register(name string, f func(*state, command) error) {
	c.cmds[name] = f
}

// runs a given command with the provided state, if it exists
func (c *commands) run(s *state, cmd command) error {

	return nil
}

func handlerLogin(s *state, cmd command) error {
	//login handler expects a single arg - the Username
	if (len(cmd.args) == 0) || (len(cmd.args) > 1) {
		return errors.New("error, expecting only username argument")
	}

	err := s.cfgPointer.SetUser(cmd.args[0])
	if err != nil {
		return err
	}

	fmt.Println("User has been set to: %s", cmd.args[0])

	return nil
}
