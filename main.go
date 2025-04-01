package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/DavidLSaldana/gator/internal/cli"
	"github.com/DavidLSaldana/gator/internal/config"
	"github.com/DavidLSaldana/gator/internal/database"
)

import _ "github.com/lib/pq"

func main() {

	//config.TestGetConfigPath()

	//read config file
	cfg, err := config.Read()
	if err != nil {
		log.Fatalln(err)
	}

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatalln(err)
	}

	dbQueries := database.New(db)

	currentState := &cli.State{
		Db:         dbQueries,
		CfgPointer: &cfg,
	}

	cmdList := make(map[string]func(*cli.State, cli.Command) error)

	commands := cli.Commands{
		Cmds: cmdList,
	}

	commands.Register("login", cli.HandlerLogin)
	commands.Register("register", cli.HandlerRegister)
	commands.Register("reset", cli.HandlerReset)
	commands.Register("users", cli.HandlerUsers)

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
