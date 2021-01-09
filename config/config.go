package config

import (
	"encoding/json"
	"io/ioutil"
	"os/user"
)

// Config -
type Config struct {
	Sources  []string  `json:"sources"`
	Packages []Package `json:"packages"`
}

// Package -
type Package struct {
	Skip    bool   `json:"skip"`
	Name    string `json:"name"`
	Flags   string `json:"flags"`
	Source  string `json:"source"`
	Version string `json:"version"`
	Command string `json:"command"`
}

// Read -
func Read() (*Config, error) {
	home, err := homeDirectory()
	if err != nil {
		return &Config{}, err
	}

	jsonFile, err := ioutil.ReadFile(home + "/.restoros/config.json")
	if err != nil {
		return &Config{}, err
	}

	cfg := &Config{}
	if err = json.Unmarshal(jsonFile, cfg); err != nil {
		return &Config{}, err
	}

	return cfg, nil
}

func homeDirectory() (string, error) {
	user, err := user.Current()
	if err != nil {
		return "", err
	}
	return user.HomeDir, nil
}
