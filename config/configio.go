package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os/user"
	"restoros/models"
)

// Read -
func Read() (*models.Config, error) {
	home, err := homeDirectory()
	if err != nil {
		return &models.Config{}, err
	}

	jsonFile, err := ioutil.ReadFile(home + "/.restoros/config.json")
	if err != nil {
		return &models.Config{}, err
	}

	cfg := &models.Config{}
	if err = json.Unmarshal(jsonFile, cfg); err != nil {
		return &models.Config{}, err
	}

	return cfg, nil
}

// Write -
func Write(*models.Config) error {
	return fmt.Errorf("Still to implement write config")
}

func homeDirectory() (string, error) {
	user, err := user.Current()
	if err != nil {
		return "", err
	}
	return user.HomeDir, nil
}
