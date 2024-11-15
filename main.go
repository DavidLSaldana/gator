package main

import (
	"fmt"
	"github.com/DavidLSaldana/gator/internal/config"
)

func main() {

	cfg := config.Config{}

	//error in read
	cfg, err := config.Read()
	if err != nil {
		return
	}

	cfg.SetUser("David")

	cfg, err = config.Read()
	if err != nil {
		return
	}

	fmt.Println(cfg.DBURL)
	fmt.Println(cfg.CurrentUserName)

}
