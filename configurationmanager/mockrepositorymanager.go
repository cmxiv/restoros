package configurationmanager

type MockRepositoryManager struct {
	InitializeReturn    error
	GetOriginReturn     string
	SetOriginCalledWith string
}

func (manager *MockRepositoryManager) Initialize() error {
	return manager.InitializeReturn
}

func (manager *MockRepositoryManager) Sync() error {
	return nil
}

func (manager *MockRepositoryManager) GetOrigin() string {
	return manager.GetOriginReturn
}

func (manager *MockRepositoryManager) SetOrigin(originUri string) error {
	manager.SetOriginCalledWith = originUri
	return nil
}
