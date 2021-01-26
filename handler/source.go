package handler

import "restoros/models"

// SourceHandler -
type SourceHandler struct {
	command *models.Command
}

// Handle -
func (sourceHandler *SourceHandler) Handle(config *models.Config) (*models.Config, error) {
	return config, nil
}
