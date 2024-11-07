package config

import (
	"encoding/json"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	Db_url   string `json:"db_url"`
	Username string `json:"current_user_name"`
}

func Read() Config {
	homeDirectory, _ := os.UserHomeDir()
	dat, _ := os.ReadFile(homeDirectory + "/.gatorconfig.json")
	cfg := Config{}

	json.Unmarshal(dat, &cfg)

	return cfg
}

// in progress
func (cfg Config) SetUser() {
	os.WriteFile()
}

func getConfigFilePath() (string, error) {
	homeDirectory, err := os.UserHomeDir()
	return homeDirectory, err
}
