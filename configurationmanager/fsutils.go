package configurationmanager

import (
	"os"
	"os/user"
	"strings"
)

const RESTOROS_DIR_NAME = ".restoros"

const CONFIGURATION_FILE_NAME = "configuration.json"

func RestorosDirectory() string {
	return homeDirectory() + string(os.PathSeparator) + RESTOROS_DIR_NAME
}

func pathFromRestorosDirectory(relativePathFromHome []string) string {
	relativePath := strings.Join(relativePathFromHome, string(os.PathSeparator))
	return RestorosDirectory() + string(os.PathSeparator) + relativePath
}

func homeDirectory() string {
	user, _ := user.Current()
	return user.HomeDir
}

func isConfigDirectoryInitialized() bool {
	_, err := os.Stat(RestorosDirectory())
	return !os.IsNotExist(err)
}
