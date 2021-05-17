package confighandler

import "restoros/configurationmanager"

type ConfigSyncHandler struct {
	RepoManager configurationmanager.RepositoryManager
}

func (handler *ConfigSyncHandler) Handle(args []string) error {
	return handler.RepoManager.Sync()
}