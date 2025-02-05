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

// config file name:
const configFileName string = ".gatorconfig.json"

// helper function for:
// getting path to json/db file
func getConfigFilePath() (string, error) {

	return "", nil
}

//helper function for:
//writing to json/db file
