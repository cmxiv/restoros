package configurationmanager

import (
	"fmt"
	"net/url"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
)

type IRepositoryManager interface {
	GetOrigin() string
	SetOrigin(string) error
}

type RepositoryManager struct {
	Path string
}

func (manager *RepositoryManager) GetOrigin() string {
	repo := manager.getRepository()
	if repo == nil {
		return ""
	}

	remote, err := repo.Remote("origin")
	if err != nil {
		return ""
	}

	return remote.Config().URLs[0]
}

func (manager *RepositoryManager) SetOrigin(originUri string) error {
	origin, err := url.ParseRequestURI(originUri)
	if err != nil || (origin != nil && origin.Host != "github.com") {
		return fmt.Errorf("invalid origin url %s (only support github urls)", originUri)
	}

	repo := manager.getRepository()
	if repo == nil {
		return fmt.Errorf("repository not found or initialized")
	}

	_, err = repo.CreateRemote(&config.RemoteConfig{
		Name: "origin",
		URLs: []string{originUri},
	})
	if err != nil {
		return fmt.Errorf("unable to set origin to %s", originUri)
	}

	return nil
}

func (manager *RepositoryManager) getRepository() *git.Repository {
	repo, err := git.PlainOpen(manager.Path)
	if err != nil {
		return nil
	}
	return repo
}
