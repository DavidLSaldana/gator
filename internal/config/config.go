package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() Config {

	return Config{}
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
