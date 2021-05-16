package configurationmanager

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

const configurationFilename = ".restoros/configuration.json"

type IManager interface {
	IsConfigurationInitialized() bool
	Write(*RestorosConfiguration) error
	Read() (*RestorosConfiguration, error)
}

type Manager struct{}

func (manager Manager) IsConfigurationInitialized() bool {
	home := homeDirectory()
	_, err := os.Stat(home + "/" + configurationFilename)
	return !os.IsNotExist(err)
}

func (manager Manager) Read() (*RestorosConfiguration, error) {
	home := homeDirectory()
	jsonFile, err := ioutil.ReadFile(home + "/" + configurationFilename)
	if err != nil {
		return nil, err
	}

	configuration := &RestorosConfiguration{}
	if err = json.Unmarshal(jsonFile, configuration); err != nil {
		return nil, err
	}

	return configuration, nil
}

func (manager Manager) Write(configuration *RestorosConfiguration) error {
	home := homeDirectory()

	if !isConfigDirectoryInitialized() {
		cwd, _ := os.Getwd()
		os.Chdir(home)
		os.Mkdir(".restoros", 0755)
		os.Chdir(cwd)
	}

	cfgByte, _ := json.Marshal(configuration)
	if err := ioutil.WriteFile(home+"/"+configurationFilename, cfgByte, 0755); err != nil {
		return err
	}

	return nil
}
