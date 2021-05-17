package configurationmanager

type MockManager struct {
	IsInitialized    bool
	ReadReturnConfig RestorosConfiguration
	ReadReturnError  error
	WriteReturnError error
}

func (manager MockManager) IsConfigurationInitialized() bool {
	return manager.IsInitialized
}

func (manager MockManager) Read() (*RestorosConfiguration, error) {
	return &manager.ReadReturnConfig, manager.ReadReturnError
}

func (manager MockManager) Write(configuration *RestorosConfiguration) error {
	return manager.WriteReturnError
}

func (manager MockManager) AddPackage(pkg Package) error {
	return nil
}
