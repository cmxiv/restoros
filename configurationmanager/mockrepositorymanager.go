package configurationmanager

type MockRepositoryManager struct {
	GetOriginReturn     string
	SetOriginCalledWith string
}

func (manager *MockRepositoryManager) GetOrigin() string {
	return manager.GetOriginReturn
}

func (manager *MockRepositoryManager) SetOrigin(originUri string) error {
	manager.SetOriginCalledWith = originUri
	return nil
}
