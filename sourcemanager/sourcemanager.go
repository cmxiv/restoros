package sourcemanager

import "restoros/sourcemanager/sources"

type Source interface {
	Find() error
	Name() string
	Purge() error
	Remove() error
	Update() error
	Install() error
}

type SourceManager interface {
	FindByPackage(string) []Source
	FindByName(string) Source
}

type sourceManager struct {
	sources []Source
}

func NewSourceManager() sourceManager {
	aptSource := sources.NewAptSource()
	brewSource := sources.NewBrewSource()
	sources := []Source{&aptSource, &brewSource}
	return sourceManager{sources: sources}
}

func (manager *sourceManager) FindByPackage(packageName string) []Source {
	return manager.sources
}

func (manager *sourceManager) FindByName(sourceName string) Source {
	for _, source := range manager.sources {
		if source.Name() == sourceName {
			return source
		}
	}
	return nil
}
