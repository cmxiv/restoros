package handler

import (
	"bufio"
	"fmt"
	"os"
	"restoros/config"
	"restoros/models"
	"strings"

	"github.com/go-git/go-git/v5"
)

// ConfigHandler -
type ConfigHandler struct {
	command *models.Command
}

// Handle -
func (configHandler *ConfigHandler) Handle(config *models.Config) (*models.Config, error) {
	switch configHandler.command.Secondary {
	case "init":
		return configHandler.init()
	case "sync":
		return configHandler.sync(config)
	case "origin":
		return configHandler.origin(config, false)
	}
	return config, nil
}

func (configHandler *ConfigHandler) init() (*models.Config, error) {
	if config.IsInitialized() {
		return config.Read(nil)
	}

	repository, err := config.CreateConfigDirectory()
	if err != nil {
		fmt.Println("Unable to create config directory")
		config.DeleteConfigDirectory()
		return &models.Config{}, err
	}

	cfg := &models.Config{
		Sources:    []string{},
		Repository: repository,
		Packages:   []models.Package{},
	}

	cfg, _ = configHandler.origin(cfg, true)
	config.Write(cfg)

	return configHandler.sync(cfg)
}

func (configHandler *ConfigHandler) sync(cfg *models.Config) (*models.Config, error) {
	worktree, _ := cfg.Repository.Worktree()
	worktree.Add(config.ConfigFileName)
	worktree.Commit("Syncing config file", &git.CommitOptions{})
	if _, err := cfg.Repository.Remote("origin"); err != nil {
		return cfg, nil
	}
	err := cfg.Repository.Push(&git.PushOptions{RemoteName: "origin"})
	fmt.Println("Push error", err)
	return cfg, nil
}

func (configHandler *ConfigHandler) origin(cfg *models.Config, userInputRequired bool) (*models.Config, error) {

	var origin string

	if userInputRequired {
		origin = askOriginFromUser()
	} else {
		origin = configHandler.command.Arguments[0]
	}

	if origin != "" {
		config.SetOrigin(origin, cfg.Repository)
		cfg.PushUpdate = true
	}

	return cfg, nil
}

func askOriginFromUser() string {
	fmt.Print("Provide origin if you want to set it (Hit enter to skip): ")
	reader := bufio.NewReader(os.Stdin)
	origin, _ := reader.ReadString('\n')
	return strings.TrimSpace(origin)
}
