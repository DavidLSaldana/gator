package main

import (
	"fmt"
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
	fmt.Printf("Initial Read: %+v\n", cfg)

	err = cfg.SetUser(currentUser)
	if err != nil {
		fmt.Println(err)
		return
	}

	cfg, err = config.Read()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Post SetUser Read: %+v\n", cfg)

}
