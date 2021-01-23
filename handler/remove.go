package handler

import "restoros/models"

// RemoveHandler -
type RemoveHandler struct {
	command *models.Command
}

// Handle -
func (removeHandler *RemoveHandler) Handle(config *models.Config) *models.Config {
	return config
}
