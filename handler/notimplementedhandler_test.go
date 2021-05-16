package handler

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNotImplementedHandler(t *testing.T) {
	handler := &NotImplementedHandler{}
	assert.Equal(t, fmt.Errorf("command not implemented, yet!"), handler.Handle(nil))
}
