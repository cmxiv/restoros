package handler

import "restoros/models"

// UpdateHandler -
type UpdateHandler struct {
	command *models.Command
}

// Handle -
func (updateHandler *UpdateHandler) Handle(config *models.Config) *models.Config {
	return config
}
