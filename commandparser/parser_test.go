package commandparser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGivenValidCLIArgumentsWhenParseCalledThenShouldReturnValidAndHandler(t *testing.T) {
	args := []string{"config", "origin", "some-link-to-origin"}
	cmd, valid := Parse(args)
	assert.True(t, valid)
	assert.NotNil(t, cmd.handler)
	assert.Equal(t, args[2:], cmd.args)
}

func TestGivenInvalidCLIArugmentsWhenParseCalledThenShouldReturnInvalid(t *testing.T) {
	cmd, valid := Parse([]string{"invalid-arg"})
	assert.Nil(t, cmd)
	assert.False(t, valid)
}
