package cli

import (
	"errors"
	"github.com/DavidLSaldana/gator/internal/state"
)

type command struct {
	name string
	args []string
}

func handlerLogin(s *state.State, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("Please enter an argument with this command")
	}
	return nil
}
