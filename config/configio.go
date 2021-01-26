package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"os/user"
	"restoros/models"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
)

// ConfigFileName -
const ConfigFileName = "config.json"

// Read -
func Read(command *models.Command) (*models.Config, error) {

	if command != nil && command.Primary == "config" && command.Secondary == "init" {
		return nil, nil
	}

	home, err := homeDirectory()
	if err != nil {
		return &models.Config{}, err
	}

	jsonFile, err := ioutil.ReadFile(configFile(home))
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
func Write(cfg *models.Config) error {
	home, err := homeDirectory()
	if err != nil {
		return err
	}

	cfgByte, err := json.Marshal(cfg)
	if err = ioutil.WriteFile(configFile(home), cfgByte, 0755); err != nil {
		return err
	}

	return nil
}

// CreateConfigDirectory -
func CreateConfigDirectory() (*git.Repository, error) {
	cwd, err := os.Getwd()
	home, err := homeDirectory()
	if err != nil {
		return nil, err
	}
	os.Chdir(home)
	os.Mkdir(".restoros", 0755)
	os.Chdir(cwd)

	repository, err := initGitRepo()

	return repository, err
}

// DeleteConfigDirectory -
func DeleteConfigDirectory() {
	home, _ := homeDirectory()
	os.RemoveAll(restorosDir(home))
}

// IsInitialized -
func IsInitialized() bool {
	_, err := Read(nil)
	return err == nil
}

// SetOrigin -
func SetOrigin(origin string, repository *git.Repository) {
	originURI, err := url.ParseRequestURI(origin)

	if err != nil || (originURI != nil && originURI.Host != "github.com") {
		fmt.Println("Invalid origin URI (ex. https://github.com/cmxiv/restoros); Unable to set origin")
		return
	}

	fmt.Println("Setting restoros origin to", origin)
	repository.CreateRemote(&config.RemoteConfig{
		Name: "origin",
		URLs: []string{origin},
	})

}

func initGitRepo() (*git.Repository, error) {
	home, err := homeDirectory()
	if err != nil {
		return nil, err
	}

	repository, err := git.PlainInit(restorosDir(home), false)
	if err != nil {
		return nil, err
	}

	return repository, nil
}

func homeDirectory() (string, error) {
	user, err := user.Current()
	if err != nil {
		return "", err
	}
	return user.HomeDir, nil
}

func configFile(home string) string {
	return restorosDir(home) +"/"+ ConfigFileName
}

func restorosDir(home string) string {
	return home + "/.restoros"
}
