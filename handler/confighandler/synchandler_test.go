package confighandler

import (
	"restoros/configurationmanager"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGivenRestorosDirectoryWhenConfigSyncCalledThenShouldSyncWithRemoteRepository(t *testing.T) {
	mock := &configurationmanager.MockRepositoryManager{}
	handler := &ConfigSyncHandler{RepoManager: mock}
	handler.Handle([]string{})
	assert.True(t, mock.SyncCalled)
}
