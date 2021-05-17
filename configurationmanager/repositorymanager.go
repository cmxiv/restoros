package configurationmanager

import (
	"fmt"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
)

type RepositoryManager interface {
	Sync() error
	Initialize() error
	GetOrigin() string
	SetOrigin(string) error
}

type repositoryManager struct {
	restorosDirectory string
}

func NewRepositoryManager(homeDirectory string) repositoryManager {
	return repositoryManager{restorosDirectory: restorosDirectory(homeDirectory)}
}

func (manager *repositoryManager) Initialize() error {
	_, err := git.PlainInit(manager.restorosDirectory, false)
	return err
}

func (manager *repositoryManager) Sync() error {
	repo := manager.getRepository()
	if repo == nil {
		return fmt.Errorf("repository doesn't exist")
	}

	worktree, _ := repo.Worktree()
	if status, _ := worktree.Status(); !status.IsClean() {
		worktree.Add(CONFIGURATION_FILE_NAME)
		worktree.Commit("sync configuration file", &git.CommitOptions{})
	}

	if _, err := repo.Remote("origin"); err != nil {
		return fmt.Errorf("no origin available. set an origin first and try to sync again")
	}

	if err := repo.Push(&git.PushOptions{RemoteName: "origin"}); err != nil {
		return fmt.Errorf("unable to push changes to remote")
	}

	return nil
}

func (manager *repositoryManager) GetOrigin() string {
	repo := manager.getRepository()
	if repo == nil {
		return "repository doesn't exist"
	}

	remote, err := repo.Remote("origin")
	if err != nil {
		return ""
	}

	return remote.Config().URLs[0]
}

func (manager *repositoryManager) SetOrigin(originUri string) error {

	if !strings.HasPrefix(originUri, "git@github.com") || !strings.HasSuffix(originUri, ".git") {
		return fmt.Errorf("invalid origin %s (only support github via ssh)", originUri)
	}

	repo := manager.getRepository()
	if repo == nil {
		return fmt.Errorf("repository not found or initialized")
	}

	_, err := repo.CreateRemote(&config.RemoteConfig{
		Name: "origin",
		URLs: []string{originUri},
	})
	if err != nil {
		return fmt.Errorf("unable to set origin to %s", originUri)
	}

	return nil
}

func (manager *repositoryManager) getRepository() *git.Repository {
	repo, err := git.PlainOpen(manager.restorosDirectory)
	if err != nil {
		return nil
	}
	return repo
}
