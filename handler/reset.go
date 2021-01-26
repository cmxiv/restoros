package handler

import "restoros/models"

// ResetHandler -
type ResetHandler struct {
	command *models.Command
}

// Handle -
func (resetHandler *ResetHandler) Handle(config *models.Config) (*models.Config, error) {
	return config, nil
}
