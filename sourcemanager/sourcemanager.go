package sourcemanager

import "restoros/sourcemanager/sources"

type Source interface {
	Find() error
	Name() string
	Purge() error
	Remove() error
	Update() error
	Install() error
	SetPackage(string, string)
}

type SourceManager interface {
	FindByPackage() []Source
	FindByName(string) Source
	SetSearchPackage(string, string)
}

type sourceManager struct {
	null Source
	sources []Source
	packageName string
	packageVersion string
}

func NewSourceManager() sourceManager {
	apt := sources.NewAptSource()
	brew := sources.NewBrewSource()
	null := sources.NewNullSource()
	sources := []Source{&apt, &brew}
	return sourceManager{sources: sources, null: &null}
}

func (manager *sourceManager) SetSearchPackage(name string, version string) {
	manager.packageName = name
	manager.packageVersion = version
}

func (manager *sourceManager) FindByPackage() []Source {
	filteredSources := []Source{}
	for _, source := range manager.sources {
		source.SetPackage(manager.packageName, manager.packageVersion)
		if err := source.Find(); err == nil {
			filteredSources = append(filteredSources, source)
		}
	}
	return filteredSources
}

func (manager *sourceManager) FindByName(sourceName string) Source {
	for _, source := range manager.sources {
		if source.Name() == sourceName {
			source.SetPackage(manager.packageName, manager.packageVersion)
			return source
		}
	}
	return manager.null
}
