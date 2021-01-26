package handler

import "restoros/models"

// PurgeHandler -
type PurgeHandler struct {
	command *models.Command
}

// Handle -
func (purgeHandler *PurgeHandler) Handle(config *models.Config) (*models.Config, error) {
	return config, nil
}
