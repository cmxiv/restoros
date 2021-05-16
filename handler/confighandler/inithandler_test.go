package confighandler

import (
	"fmt"
	"restoros/configurationmanager"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGivenRestorosAlreadyInitializedWhenConfigInitCalledThenShouldReturnAlreadyInitializedError(t *testing.T) {
	handler := setupInitHandler(true, nil)
	assert.EqualError(t, handler.Handle([]string{}), "already initialized")
}

func TestGivenRestorosNotInitializedWhenConfigInitCalledThenShouldCallConfigurationManagerWrite(t *testing.T) {
	handler := setupInitHandler(false, fmt.Errorf("configmanager write called"))
	assert.EqualError(t, handler.Handle([]string{}), "configmanager write called")
}

func setupInitHandler(isInitialized bool, writeError error) *ConfigInitHandler {
	return &ConfigInitHandler{
		Manager: configurationmanager.MockManager{IsInitialized: isInitialized, WriteReturnError: writeError},
	}
}
