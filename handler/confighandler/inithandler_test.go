package confighandler

import (
	"fmt"
	"restoros/configurationmanager"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGivenRestorosAlreadyInitializedWhenConfigInitCalledThenShouldReturnAlreadyInitializedError(t *testing.T) {
	handler := setupInitHandler(true, nil, nil)
	assert.EqualError(t, handler.Handle([]string{}), "already initialized")
}

func TestGivenRestorosNotInitializedWhenConfigInitCalledThenShouldCallConfigurationManagerWrite(t *testing.T) {
	handler := setupInitHandler(false, nil, fmt.Errorf("configmanager write called"))
	assert.EqualError(t, handler.Handle([]string{}), "configmanager write called")
}

func TestWhenErrorWhileCreatingRepositoryThenShouldReturnError(t *testing.T) {
	handler := setupInitHandler(false, fmt.Errorf("error while initialization"), nil)
	assert.EqualError(t, handler.Handle([]string{}), "error while initialization")
}

func setupInitHandler(isInitialized bool, repoError error, writeError error) *ConfigInitHandler {
	return &ConfigInitHandler{
		Manager:     &configurationmanager.MockManager{IsInitialized: isInitialized, WriteReturnError: writeError},
		RepoManager: &configurationmanager.MockRepositoryManager{InitializeReturn: repoError},
	}
}
