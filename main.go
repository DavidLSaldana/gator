package main

import (
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

	cmdList := make(map[string]func(*cli.State, cli.Command) error)

	commands := cli.Commands{
		Cmds: cmdList,
	}

	commands.Register("login", cli.HandlerLogin)

	args := os.Args
	if len(args) < 2 {
		log.Fatalln("Error: not enough arguments provided")
	}

	commandName := args[1]
	commandArgs := args[2:]

	command := cli.Command{
		Name: commandName,
		Args: commandArgs,
	}

	err = commands.Run(currentState, command)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Current State: %+v\n", cfg)

}
