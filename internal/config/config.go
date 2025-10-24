package config

import (
	"encoding/json"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	Db_url            string `json:"db_url"`
	Current_user_name string `json:"current_user_name"`
}

func (c *Config) SetUser(username string) error {
	c.Current_user_name = username
	return write(*c)
}

func Read() (Config, error) {
	filePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	byteData, err := os.ReadFile(filePath)
	if err != nil {
		return Config{}, err
	}

	configData := Config{}
	err = json.Unmarshal(byteData, &configData)
	if err != nil {
		return Config{}, err
	}

	return configData, nil
}

func getConfigFilePath() (string, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	filePath := homePath + "/" + configFileName
	return filePath, nil
}

func write(config Config) error {
	filePath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	f, err := os.Create(filePath)
	if err != nil {
		return err
	}

	defer f.Close()

	encoder := json.NewEncoder(f)
	if err := encoder.Encode(config); err != nil {
		return err
	}

	return nil
}
