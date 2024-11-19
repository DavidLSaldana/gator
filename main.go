package main

import (
	"fmt"
	"github.com/DavidLSaldana/gator/internal/cli"
	"github.com/DavidLSaldana/gator/internal/config"
	"github.com/DavidLSaldana/gator/internal/state"
)

func main() {

	//error in read
	cfg, err := config.Read()
	if err != nil {
		return
	}

	mainState := state.State{
		Cfg: &cfg,
	}

	cliMap := make(map[string]func(*state.State, cli.Command) error)

	commands := cli.Commands{
		CommandMap: cliMap,
	}

	//I think I'm doing something wrong here, but not totally sure yet. I'll
	// keep going and see if I come across the right way to do this.
	commands.Register("login", commandLogin(mainState, cli.Command) {

	})
}
