package confighandler

import (
	"io/ioutil"
	"os"
	"restoros/configurationmanager"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhenConfigOriginCalledWithoutArgumentThenShouldOutputCurrentSetOrigin(t *testing.T) {
	handler, _, consoleOutput, reset := setupOriginHandler(true, "github-repository-url")
	err := handler.Handle([]string{})
	reset()
	assert.Nil(t, err)
	assert.Equal(t, "github-repository-url", strings.TrimSpace(consoleOutput()))
}

func TestWhenOriginCalledWithArgumentThenSetOriginShouldBeCalledWithFirstArgument(t *testing.T) {
	handler, mockRepoManager, _, _ := setupOriginHandler(false, "")
	handler.Handle([]string{"github-repository-url", "ignore-this-argument"})
	assert.Equal(t, "github-repository-url", mockRepoManager.SetOriginCalledWith)
}

func setupOriginHandler(spyOnStdout bool, getOriginReturn string) (*ConfigOriginHandler, *configurationmanager.MockRepositoryManager, func() string, func()) {

	var (
		r              *os.File
		w              *os.File
		preserveStdout *os.File
	)

	if spyOnStdout {
		preserveStdout = os.Stdout
		r, w, _ = os.Pipe()
		os.Stdout = w
	}

	consoleOutput := func() string {
		out, _ := ioutil.ReadAll(r)
		return string(out)
	}

	reset := func() {
		w.Close()
		os.Stdout = preserveStdout
	}

	mock := &configurationmanager.MockRepositoryManager{
		GetOriginReturn: getOriginReturn,
	}

	handler := &ConfigOriginHandler{
		RepoManager: mock,
	}

	return handler, mock, consoleOutput, reset
}
