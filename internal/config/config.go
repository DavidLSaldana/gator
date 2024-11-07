package config

import (
	"os"
)

type Config struct {
	db_url            string `db_url: "connection_string_goes_here"`
	current_user_name string `current_user_name: "username_goes_here"`
}

func Read() Config {
	//in progress
	file_location, _ := os.UserHomeDir()
	os.Open(file_location + ".gatorconfig.json")
	return Config{}
}
