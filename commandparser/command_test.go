package commandparser

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockHandler struct{}

func (mockHandler *MockHandler) Handle(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("mock handler")
	}
	return nil
}

func TestGivenCommandWhenExecCalledThenShouldInvokeHandlerAndReturnNil(t *testing.T) {
	handler := &MockHandler{}
	command := Command{handler: handler, args: []string{"foo", "bar"}}
	assert.Nil(t, command.Exec())
}

func TestGivenCommandWhenExecCalledThenShouldInvokeHandlerAndReturnError(t *testing.T) {
	handler := &MockHandler{}
	command := Command{handler: handler}
	assert.Error(t, command.Exec())
}
