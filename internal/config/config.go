package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	//get file path
	path, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	//open file at path
	file, err := os.Open(path)
	if err != nil {
		return Config{}, err
	}

	//defer to ensure file closes
	defer file.Close()

	//decode file into variable
	//if error return Config{}, error
	decoder := json.NewDecoder(file)
	cfg := Config{}
	err = decoder.Decode(&cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func (cfg Config) SetUser() error {
	err := write(cfg)
	if err != nil {
		return err
	}

	//succussful write
	return nil
}

// config file name:
const configFileName string = ".gatorconfig.json"

// checking that getConfigFilePath() doesn't get extra characters or have
// missing characters can REMOVE
func TestGetConfigPath() {
	homePath, err := getConfigFilePath()
	if err != nil {
		fmt.Println("ERROR")
		return
	}
	fmt.Println(homePath)
}

// helper function for:
// getting path to json/db file
func getConfigFilePath() (string, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	fullPath := homePath + "/" + configFileName

	return fullPath, nil
}

// helper function for:
// writing to json/db file
func write(cfg Config) error {
	return nil
}
