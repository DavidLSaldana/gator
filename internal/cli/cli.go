package cli

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/DavidLSaldana/gator/internal/config"
	"github.com/DavidLSaldana/gator/internal/database"
	"github.com/google/uuid"
)

type State struct {
	Db         *database.Queries
	CfgPointer *config.Config
}

type Command struct {
	Name string
	Args []string
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
	function, ok := c.Cmds[cmd.Name]
	if !ok {
		return errors.New("cmd not found")
	}

	return function(s, cmd)

}

func HandlerLogin(s *State, cmd Command) error {
	//login handler expects a single arg - the Username
	if len(cmd.Args) > 1 {
		return errors.New("error, expecting only username argument")
	}

	if len(cmd.Args) == 0 {
		return errors.New("error, a username is required")
	}

	_, err := s.Db.GetUser(context.Background(), cmd.Args[0])
	if err != nil {
		os.Exit(1)
	}

	err = s.CfgPointer.SetUser(cmd.Args[0])
	if err != nil {
		return err
	}

	fmt.Printf("User has been set to: %s\n", cmd.Args[0])

	return nil
}

func HandlerRegister(s *State, cmd Command) error {
	if len(cmd.Args) > 1 {
		return errors.New("error, expecting only username argument")
	}

	if len(cmd.Args) == 0 {
		return errors.New("error, a username is required")
	}

	newID := uuid.New()

	args := database.CreateUserParams{
		ID:        int32(newID.ID()),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0],
	}
	fmt.Println("The code makes it here!")
	newUser, err := s.Db.CreateUser(context.Background(), args)
	if err != nil {
		os.Exit(1)
	}

	s.CfgPointer.CurrentUserName = newUser.Name
	fmt.Printf("New User: %s, has been created!\n", s.CfgPointer.CurrentUserName)
	return nil
}
