package confighandler

import (
	"fmt"
	"restoros/configurationmanager"
)

type ConfigOriginHandler struct {
	RepoManager configurationmanager.IRepositoryManager
}

func (handler *ConfigOriginHandler) Handle(args []string) error {
	if len(args) < 1 {
		fmt.Println(handler.RepoManager.GetOrigin())
		return nil
	}
	return handler.RepoManager.SetOrigin(args[0])
}
