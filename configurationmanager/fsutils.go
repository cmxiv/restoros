package configurationmanager

import (
	"os"
	"os/user"
)

func homeDirectory() string {
	user, _ := user.Current()
	return user.HomeDir
}

func isConfigDirectoryInitialized() bool {
	_, err := os.Stat(RestorosDirectory())
	return !os.IsNotExist(err)
}

func RestorosDirectory() string {
	return homeDirectory() + "/.restoros"
}