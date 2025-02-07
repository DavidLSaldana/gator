package cli

import (
	"errors"
	"fmt"
	"github.com/DavidLSaldana/gator/internal/config"
)

type State struct {
	CfgPointer *config.Config
}

type Command struct {
	name string
	args []string
}

type Commands struct {
	Cmds map[string]func(*State, Command) error
}

// registers a new handler function for a Command name
func (c *Commands) Register(name string, f func(*State, Command) error) {
	c.Cmds[name] = f
}

// runs a given Command with the provided state, if it exists
func (c *Commands) Run(s *State, cmd Command) error {
	function, ok := c.Cmds[cmd.name]
	if !ok {
		return errors.New("cmd not found")
	}

	return function(s, cmd)

}

func HandlerLogin(s *State, cmd Command) error {
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
