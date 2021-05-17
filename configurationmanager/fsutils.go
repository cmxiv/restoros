package configurationmanager

import (
	"os"
	"os/user"
)

const RESTOROS_DIR_NAME = ".restoros"

const CONFIGURATION_FILE_NAME = "configuration.json"

func HomeDirectory() string {
	user, _ := user.Current()
	return user.HomeDir
}

func restorosDirectory(homeDirectory string) string {
	return homeDirectory + string(os.PathSeparator) + RESTOROS_DIR_NAME
}

func configurationFileName(homeDirectory string) string {
	return restorosDirectory(homeDirectory) + string(os.PathSeparator) + CONFIGURATION_FILE_NAME
}

func isDirectoryInitialized(restorosDirectory string) bool {
	_, err := os.Stat(restorosDirectory)
	return !os.IsNotExist(err)
}
