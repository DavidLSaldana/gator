package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/DavidLSaldana/gator/internal/cli"
	"github.com/DavidLSaldana/gator/internal/config"
	"github.com/DavidLSaldana/gator/internal/database"
)

import _ "github.com/lib/pq"

func main() {

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
	commands.Register("agg", cli.HandlerAgg)
	commands.Register("addfeed", cli.MiddlewareLoggedIn(cli.HandlerAddFeed))
	commands.Register("feeds", cli.HandlerFeeds)
	commands.Register("follow", cli.MiddlewareLoggedIn(cli.HandlerFollow))
	commands.Register("following", cli.MiddlewareLoggedIn(cli.HandlerFollowing))

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

}
