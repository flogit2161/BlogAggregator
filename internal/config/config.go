package config

import (
	"encoding/json"
	"errors"
	"os"
)

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (c *Config) SetUser(name string) error {
	c.CurrentUserName = name
	path, err := getConfigFilePath()
	if err != nil {
		return errors.New("Could not get filepath")
	}

	data, err := json.Marshal(c)
	if err != nil {
		return errors.New("Could not Marshal the config")
	}

	err = os.WriteFile(path, data, 0o600)
	if err != nil {
		return errors.New("Could not write to file")
	}

	return err
}

func ReadConfig() (Config, error) {
	filepath, err := getConfigFilePath()
	if err != nil {
		return Config{}, errors.New("Could not get filepath")
	}

	data, err := os.ReadFile(filepath)
	if err != nil {
		return Config{}, errors.New("Could not read file from filepath")
	}

	cfg := Config{}
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return Config{}, errors.New("Could not Unmarshal file")
	}
	return cfg, nil

}

func getConfigFilePath() (string, error) {
	path, err := os.UserHomeDir()
	if err != nil {
		return "", errors.New("Could not get user's home path")
	}
	filepath := path + "/.gatorconfig.json"
	return filepath, nil
}
