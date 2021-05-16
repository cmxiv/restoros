package configurationmanager

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type IManager interface {
	IsConfigurationInitialized() bool
	Write(*RestorosConfiguration) error
	Read() (*RestorosConfiguration, error)
}

type Manager struct{}

func (manager *Manager) IsConfigurationInitialized() bool {
	_, err := os.Stat(pathFromRestorosDirectory([]string{CONFIGURATION_FILE_NAME}))
	return !os.IsNotExist(err)
}

func (manager *Manager) Read() (*RestorosConfiguration, error) {
	jsonFile, err := ioutil.ReadFile(pathFromRestorosDirectory([]string{CONFIGURATION_FILE_NAME}))
	if err != nil {
		return nil, err
	}

	configuration := &RestorosConfiguration{}
	if err = json.Unmarshal(jsonFile, configuration); err != nil {
		return nil, err
	}

	return configuration, nil
}

func (manager *Manager) Write(configuration *RestorosConfiguration) error {
	home := homeDirectory()

	if !isConfigDirectoryInitialized() {
		cwd, _ := os.Getwd()
		os.Chdir(home)
		os.Mkdir(RESTOROS_DIR_NAME, 0755)
		os.Chdir(cwd)
	}

	cfgByte, _ := json.Marshal(configuration)
	filePath := pathFromRestorosDirectory([]string{CONFIGURATION_FILE_NAME})
	if err := ioutil.WriteFile(filePath, cfgByte, 0755); err != nil {
		return err
	}

	return nil
}
