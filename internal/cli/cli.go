package cli

import (
	"errors"
	"fmt"
	"github.com/DavidLSaldana/gator/internal/config"
)

type State struct {
	CfgPointer *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	cmds map[string]func(*State, command) error
}

// registers a new handler function for a command name
func (c *commands) register(name string, f func(*State, command) error) {
	c.cmds[name] = f
}

// runs a given command with the provided state, if it exists
func (c *commands) run(s *State, cmd command) error {
	function, ok := c.cmds[cmd.name]
	if !ok {
		return errors.New("cmd not found")
	}

	return function(s, cmd)

}

func handlerLogin(s *State, cmd command) error {
	//login handler expects a single arg - the Username
	if (len(cmd.args) == 0) || (len(cmd.args) > 1) {
		return errors.New("error, expecting only username argument")
	}

	err := s.CfgPointer.SetUser(cmd.args[0])
	if err != nil {
		return err
	}

	fmt.Printf("User has been set to: %s\n", cmd.args[0])

	return nil
}
