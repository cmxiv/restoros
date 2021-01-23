package handler

import "restoros/models"

// ConfigHandler -
type ConfigHandler struct {
	command *models.Command
}

// Handle -
func (configHandler *ConfigHandler) Handle(config *models.Config) *models.Config {
	return config
}
