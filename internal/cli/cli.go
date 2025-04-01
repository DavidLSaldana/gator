package cli

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/DavidLSaldana/gator/internal/config"
	"github.com/DavidLSaldana/gator/internal/database"
	"github.com/DavidLSaldana/gator/internal/rss"
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
	fmt.Printf("This has happened")

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
	err = s.CfgPointer.SetUser(newUser.Name)
	if err != nil {
		return err
	}
	return nil
}

func HandlerReset(s *State, cmd Command) error {
	if len(cmd.Args) != 0 {
		return errors.New("not expecting any arguments for 'reset' command")
	}

	err := s.Db.ResetUsers(context.Background())
	if err != nil {
		os.Exit(1)
	}

	return nil
}

func HandlerUsers(s *State, cmd Command) error {
	if len(cmd.Args) != 0 {
		return errors.New("not expecting any arguments for 'users' command")
	}

	listOfUsers, err := s.Db.GetUsers(context.Background())
	if err != nil {
		return err
	}

	for _, user := range listOfUsers {
		if user == s.CfgPointer.CurrentUserName {
			fmt.Printf("* %s (current)\n", user)
		} else {
			fmt.Printf("* %s\n", user)
		}
	}

	return nil
}

func HandlerAgg(s *State, cmd Command) error {
	feed := &rss.RSSFeed{}
	url := "https://www.wagslane.dev/index.xml"
	feed, err := rss.FetchFeed(context.Background(), url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("RSSFeed Title: %s", feed.Channel.Title)
	fmt.Printf("RSSFeed Link: %s", feed.Channel.Link)
	fmt.Printf("RSSFeed Description: %s", feed.Channel.Description)

	for _, item := range feed.Channel.Item {
		fmt.Printf("Item Title: %s", item.Title)
		fmt.Printf("Item Link: %s", item.Link)
		fmt.Printf("Item Description: %s", item.Description)
		fmt.Printf("Item Publish Date: %s", item.PubDate)
	}

	return nil
}
