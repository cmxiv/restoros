package handler

import (
	"fmt"
	"restoros/argumentparser"
	"restoros/models"
)

var (
	install Handler
	update  Handler
	remove  Handler
	purge   Handler
	restore Handler
	reset   Handler
	source  Handler
	cfg     Handler
	list    Handler
)

// Handler -
type Handler interface {
	Handle(*models.Config) (*models.Config, error)
}

// GetHandler - Singleton - Not Thread safe (Yet!)
func GetHandler(command *models.Command) (Handler, error) {
	switch command.Primary {
	case "install":
		if install == nil {
			install = &InstallHandler{command: command}
		}
		return install, nil
	case "update":
		if update == nil {
			update = &UpdateHandler{command: command}
		}
		return update, nil
	case "remove":
		if remove == nil {
			remove = &RemoveHandler{command: command}
		}
		return remove, nil
	case "purge":
		if purge == nil {
			purge = &PurgeHandler{command: command}
		}
		return purge, nil
	case "restore":
		if restore == nil {
			restore = &RestoreHandler{command: command}
		}
		return restore, nil
	case "reset":
		if reset == nil {
			reset = &ResetHandler{command: command}
		}
		return reset, nil
	case "source":
		if source == nil {
			source = &SourceHandler{command: command}
		}
		return source, nil
	case "config":
		if cfg == nil {
			cfg = &ConfigHandler{command: command}
		}
		return cfg, nil
	case "list":
		if list == nil {
			list = &ListHandler{command: command}
		}
		return list, nil
	default:
		return nil, fmt.Errorf("Invalid command" + argumentparser.UsageMessage)
	}
}
