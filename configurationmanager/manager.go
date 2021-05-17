package configurationmanager

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Manager interface {
	AddPackage(Package) error
	IsConfigurationInitialized() bool
	Write(*RestorosConfiguration) error
	Read() (*RestorosConfiguration, error)
}

type manager struct {
	homeDirectory         string
	restorosDirectory     string
	configurationFileName string
}

func NewManager(homeDirectory string) manager {
	return manager{
		homeDirectory:         homeDirectory,
		restorosDirectory:     restorosDirectory(homeDirectory),
		configurationFileName: configurationFileName(homeDirectory),
	}
}

func (manager *manager) IsConfigurationInitialized() bool {
	_, err := os.Stat(manager.configurationFileName)
	return !os.IsNotExist(err)
}

func (manager *manager) Read() (*RestorosConfiguration, error) {
	jsonFile, err := ioutil.ReadFile(manager.configurationFileName)
	if err != nil {
		return nil, err
	}

	configuration := &RestorosConfiguration{}
	if err = json.Unmarshal(jsonFile, configuration); err != nil {
		return nil, err
	}

	return configuration, nil
}

func (manager *manager) Write(configuration *RestorosConfiguration) error {

	if !isDirectoryInitialized(manager.restorosDirectory) {
		cwd, _ := os.Getwd()
		os.Chdir(manager.homeDirectory)
		os.Mkdir(RESTOROS_DIR_NAME, 0755)
		os.Chdir(cwd)
	}

	cfgByte, _ := json.Marshal(configuration)
	if err := ioutil.WriteFile(manager.configurationFileName, cfgByte, 0755); err != nil {
		return err
	}

	return nil
}

func (manager *manager) AddPackage(packageDetail Package) error {
	return fmt.Errorf("not implemented")
}
