package main

import (
	"fmt"
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

	//in progress
	fmt.Println(mainState.Cfg.DBURL)
	fmt.Println(mainState.Cfg.CurrentUserName)

}
