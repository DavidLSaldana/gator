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

	user, err := s.Db.GetUser(context.Background(), cmd.Args[0])
	if err != nil {
		os.Exit(1)
	}
	//want to update to include the user id in config struct **HERE**
	err = s.CfgPointer.SetUser(user)
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
	newUser, err := s.Db.CreateUser(context.Background(), args)
	if err != nil {
		os.Exit(1)
	}

	err = s.CfgPointer.SetUser(newUser)
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

	fmt.Printf("RSSFeed Title: %s\n", feed.Channel.Title)
	fmt.Printf("RSSFeed Link: %s\n", feed.Channel.Link)
	fmt.Printf("RSSFeed Description: %s\n", feed.Channel.Description)

	for _, item := range feed.Channel.Item {
		fmt.Printf("Item Title: %s\n", item.Title)
		fmt.Printf("Item Link: %s\n", item.Link)
		fmt.Printf("Item Description: %s\n", item.Description)
		fmt.Printf("Item Publish Date: %s\n\n", item.PubDate)
	}

	return nil
}

func HandlerAddFeed(s *State, cmd Command) error {
	if len(cmd.Args) != 2 {
		return errors.New("need a name for the feed and a URL for the feed, nothing more and nothing less")
	}
	currentTime := time.Now()
	newID := uuid.New()

	args := database.CreateFeedParams{
		ID:        int32(newID.ID()),
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		Name:      cmd.Args[0],
		Url:       cmd.Args[1],
		UserID:    s.CfgPointer.CurrentUserID,
	}
	_, err := s.Db.CreateFeed(context.Background(), args)
	if err != nil {
		return err
	}

	return nil
}

func HandlerFeeds(s *State, cmd Command) error {
	feeds, err := s.Db.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	if len(feeds) == 0 {
		return errors.New("There are no feeds to show")
	}

	for _, feed := range feeds {
		fmt.Printf("Name: %s\n", feed.Name)
		fmt.Printf("URL: %s\n", feed.Url)
		owner, err := s.Db.GetUserNameFromID(context.Background(), feed.UserID)
		if err != nil {
			return err
		}
		fmt.Printf("Owner: %s\n", owner)
	}

	return nil
}
