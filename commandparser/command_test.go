package commandparser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGivenCommandWhenExecCalledThenShouldInvokeHandler(t *testing.T) {
	handler := func(args []string) bool { return len(args) == 2 }
	command := Command{handler: handler, args: []string{"foo", "bar"}}
	assert.True(t, command.Exec())
}