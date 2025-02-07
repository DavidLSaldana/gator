package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/DavidLSaldana/gator/internal/cli"
	"github.com/DavidLSaldana/gator/internal/config"
)

func main() {

	//config.TestGetConfigPath()

	//read config file
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
		return
	}

	currentState := &cli.State{
		CfgPointer: &cfg,
	}

	commands := cli.Commands{}

	commands.Register("login", cli.HandlerLogin)

	args := os.Args
	if len(args) < 2 {
		log.Fatalln("Error: incorrect number of arguments")
	}

	commandName := args[1]
	commandArgs := args[2:]

	//this isn't working yet
	//err := commands.Run(currentState, commands.commandName)

	fmt.Printf("Post SetUser Read: %+v\n", cfg)

}
