package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/DavidLSaldana/gator/internal/database"
)

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
	CurrentUserID   int32  `json:"current_user_id"`
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

func (cfg *Config) SetUser(user database.User) error {

	cfg.CurrentUserName = user.Name
	cfg.CurrentUserID = user.ID

	err := write(*cfg)
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
	//get path to file
	path, err := getConfigFilePath()
	if err != nil {
		return err
	}

	// create/open file
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	// close file for sure, since it is now open
	defer file.Close()

	// new encoder for json into the file
	encoder := json.NewEncoder(file)
	err = encoder.Encode(cfg)
	if err != nil {
		return err
	}

	//success
	return nil
}
