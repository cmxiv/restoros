package configurationmanager

import (
	"fmt"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
)

type IRepositoryManager interface {
	Initialize() error
	Sync() error
	GetOrigin() string
	SetOrigin(string) error
}

type RepositoryManager struct {
	Path string
}

func (manager *RepositoryManager) Initialize() error {
	_, err := git.PlainInit(manager.Path, false)
	return err
}

func (manager *RepositoryManager) Sync() error {
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

func (manager *RepositoryManager) GetOrigin() string {
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

func (manager *RepositoryManager) SetOrigin(originUri string) error {

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

func (manager *RepositoryManager) getRepository() *git.Repository {
	repo, err := git.PlainOpen(manager.Path)
	if err != nil {
		return nil
	}
	return repo
}
