package config

import (
	"encoding/json"
	"os"
)

const configFileName = ".gatorconfig.json"

type config struct {
	Db_url            string
	Current_user_name string
}

func (c config) SetUser() error {
	jsonByte, err := json.Marshal(c)
	if err != nil {
		return err
	}

	err = write(jsonByte)
	if err != nil {
		return err
	}

	return nil
}

func Read() (config, error) {
	filePath, err := getConfigFilePath()
	if err != nil {
		return config{}, err
	}

	byteData, err := os.ReadFile(filePath)
	if err != nil {
		return config{}, err
	}

	configData := config{}
	err = json.Unmarshal(byteData, &configData)
	if err != nil {
		return config{}, err
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

func write(data []byte) error {
	filePath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}

	defer f.Close()

	if _, err := f.Write(data); err != nil {
		return err
	}

	return nil
}
