package configurationmanager

import (
	"fmt"
	"testing"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/stretchr/testify/assert"
)

const origin = "git@github.com:cmxiv/systemconfig.git"

func TestGivenDirectoryWithGitAndOriginSetWhenGetOriginCalledThenReturnOriginUrl(t *testing.T) {
	tmpDirPath := t.TempDir()
	setupGitWithOrigin(tmpDirPath, origin)

	manager := repositoryManager{restorosDirectory: tmpDirPath}
	assert.Equal(t, origin, manager.GetOrigin())
}

func TestGivenDirectoryWithGitInitAndNoOriginSetWhenGetOriginCalledThenReturnEmptyString(t *testing.T) {
	tmpDirPath := t.TempDir()
	setupGit(tmpDirPath)

	manager := repositoryManager{restorosDirectory: tmpDirPath}
	assert.Equal(t, "", manager.GetOrigin())
}

func TestGivenDirectoryWithoutGitInitSetWhenGetOriginCalledThenReturnConfigurationDoesntExist(t *testing.T) {
	tmpDirPath := t.TempDir()
	manager := repositoryManager{restorosDirectory: tmpDirPath}
	assert.Equal(t, "repository doesn't exist", manager.GetOrigin())
}

func TestGivenGitDirectoryWhenSetOriginCalledWithValidURLThenShouldSetOriginToProvidedURL(t *testing.T) {
	tmpDirPath := t.TempDir()
	repository := setupGit(tmpDirPath)

	manager := repositoryManager{restorosDirectory: tmpDirPath}

	if assert.Nil(t, manager.SetOrigin(origin)) {
		remote, _ := repository.Remote("origin")
		assert.Equal(t, origin, remote.Config().URLs[0])
	}
}

func TestWhenSetOriginCalledWithInvalidURLThenReturnInvalidUrlError(t *testing.T) {
	manager := repositoryManager{restorosDirectory: ""}
	assert.EqualError(t, manager.SetOrigin("foobar"), "invalid origin foobar (only support github via ssh)")
}

func TestWhenSetOriginCalledWithHostNotGithubThenReturnInvalidUrlError(t *testing.T) {
	manager := repositoryManager{restorosDirectory: ""}
	notGithubUrl := "https://foobar.com/bizbaz"
	assert.EqualError(t, manager.SetOrigin(notGithubUrl), fmt.Sprintf("invalid origin %s (only support github via ssh)", notGithubUrl))
}

func TestGivenNotGitInitializedDirectoryWhenSetOriginCalledThenReturnRepositoryNotFoundError(t *testing.T) {
	tmpDirPath := t.TempDir()
	manager := repositoryManager{restorosDirectory: tmpDirPath}
	assert.EqualError(t, manager.SetOrigin(origin), "repository not found or initialized")
}

func TestGivenEmptyDirectoryWhenInitializeCalledThenShouldInitializeGitRepository(t *testing.T) {
	tmpDirPath := t.TempDir()
	manager := repositoryManager{restorosDirectory: tmpDirPath}
	assert.Nil(t, manager.Initialize())
	assert.True(t, isGitInitialized(tmpDirPath))
}

func setupGit(tmpDirPath string) *git.Repository {
	tmpRepo, _ := git.PlainInit(tmpDirPath, false)
	return tmpRepo
}

func setupGitWithOrigin(tmpDirPath string, origin string) {
	tmpRepo := setupGit(tmpDirPath)
	tmpRepo.CreateRemote(&config.RemoteConfig{
		Name: "origin",
		URLs: []string{origin},
	})
}

func isGitInitialized(tmpDirPath string) bool {
	_, err := git.PlainOpen(tmpDirPath)
	return err == nil
}
