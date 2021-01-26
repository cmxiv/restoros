package handler

import "restoros/models"

// RestoreHandler -
type RestoreHandler struct {
	command *models.Command
}

// Handle -
func (restoreHandler *RestoreHandler) Handle(config *models.Config) (*models.Config, error) {
	return config, nil
}
