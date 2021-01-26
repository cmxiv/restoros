package handler

import (
	"restoros/models"
)

// InstallHandler -
type InstallHandler struct {
	command *models.Command
}

// Handle -
func (installHandler *InstallHandler) Handle(config *models.Config) (*models.Config, error) {
	// Multi-threaded search for the package
	// Wait untill all threads return
	// If multiple positive responses
	// 		return with the list of the responses and ask for the source of the package from the user
	// Else If all negative responses
	//		 return with negative user response
	// Else
	// 		Install the package and update the config file; Marking config as modified
	return config, nil
}
