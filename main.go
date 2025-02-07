package main

import (
	"fmt"
	"github.com/DavidLSaldana/gator/internal/cli"
	"github.com/DavidLSaldana/gator/internal/config"
)

func main() {

	//config.TestGetConfigPath()

	currentUser := "david"

	//read config file
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
		return
	}

	currentState := &cli.State{
		CfgPointer: &cfg,
	}

	fmt.Printf("Post SetUser Read: %+v\n", cfg)

}
