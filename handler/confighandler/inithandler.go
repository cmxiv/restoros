package confighandler

import (
	"fmt"
	"restoros/configurationmanager"
)

type ConfigInitHandler struct {
	Manager configurationmanager.IManager
}

func (handler *ConfigInitHandler) Handle(args []string) error {
	if handler.Manager.IsConfigurationInitialized() {
		return fmt.Errorf("already initialized")
	}
	initialConfiguration := &configurationmanager.RestorosConfiguration{}
	return handler.Manager.Write(initialConfiguration)
}