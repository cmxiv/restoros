package argumentparser

import (
	"restoros/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Test struct {
	isNegative     bool
	expectedError  string
	expectedOutput *models.Command
	input          []string
}

func TestPrintsUsageWhenNoCommandsPasses(t *testing.T) {
	_, err := Parse([]string{"restoros"})
	assert.EqualError(t, err, UsageMessage)
}

func TestPrimaryOnlyCommands(t *testing.T) {

	tests := map[string]Test{
		"Verify throws error when incorrect primary command": {
			isNegative:    true,
			expectedError: "\nInvalid Command: foo-bar " + UsageMessage,
			input:         []string{"foo-bar"},
		},
		"Verify parses `restore` command": {
			expectedOutput: &models.Command{Primary: "restore"},
			input:          []string{"restore"},
		},
		"Verify ignores any other arguments passes along with `restore`": {
			expectedOutput: &models.Command{Primary: "restore"},
			input:          []string{"restore", "fake-argument"},
		},
		"Verify parses `reset` command": {
			expectedOutput: &models.Command{Primary: "reset"},
			input:          []string{"reset"},
		},
		"Verify ignores any other arguments passes along with `reset`": {
			expectedOutput: &models.Command{Primary: "reset"},
			input:          []string{"reset", "fake-argument"},
		},
		"Verify parses `list` command": {
			expectedOutput: &models.Command{Primary: "list"},
			input:          []string{"list"},
		},
		"Verify ignores any other arguments passes along with `list`": {
			expectedOutput: &models.Command{Primary: "list"},
			input:          []string{"list", "fake-argument"},
		},
	}

	runTableTest(t, tests)

}

func TestPrimaryWithArgumentsCommands(t *testing.T) {
	tests := map[string]Test{
		"Verify throws error when only `install` command is passed": {
			isNegative:    true,
			expectedError: "\nInvalid Command: install requires package name" + UsageMessage,
			input:         []string{"install"},
		},
		"Verify parses other arguments passes along with `install`": {
			expectedOutput: &models.Command{Primary: "install", Arguments: []string{"package-name"}},
			input:          []string{"install", "package-name"},
		},
		"Verify throws error when only `update` command is passed": {
			isNegative:    true,
			expectedError: "\nInvalid Command: update requires package name" + UsageMessage,
			input:         []string{"update"},
		},
		"Verify parses other arguments passes along with `update`": {
			expectedOutput: &models.Command{Primary: "update", Arguments: []string{"package-name"}},
			input:          []string{"update", "package-name"},
		},
		"Verify throws error when only `remove` command is passed": {
			isNegative:    true,
			expectedError: "\nInvalid Command: remove requires package name" + UsageMessage,
			input:         []string{"remove"},
		},
		"Verify parses other arguments passes along with `remove`": {
			expectedOutput: &models.Command{Primary: "remove", Arguments: []string{"package-name"}},
			input:          []string{"remove", "package-name"},
		},
		"Verify throws error when only `purge` command is passed": {
			isNegative:    true,
			expectedError: "\nInvalid Command: purge requires package name" + UsageMessage,
			input:         []string{"purge"},
		},
		"Verify parses other arguments passes along with `purge`": {
			expectedOutput: &models.Command{Primary: "purge", Arguments: []string{"package-name"}},
			input:          []string{"purge", "package-name"},
		},
	}

	runTableTest(t, tests)
}

func TestPrimaryWithSecondaryArgumentsCommands(t *testing.T) {
	tests := map[string]Test{
		"Verify throws error when only `source` command is passed": {
			isNegative:    true,
			expectedError: "\nInvalid Command: source requires a sub-option" + UsageMessage,
			input:         []string{"source"},
		},
		"Verify throws error when in-correct sub-option is passed with `source`": {
			isNegative:    true,
			expectedError: "\nInvalid Command: source foo-bar" + UsageMessage,
			input:         []string{"source", "foo-bar"},
		},
		"Verify throws error when no source name is provided with `source add`": {
			isNegative:    true,
			input:         []string{"source", "add"},
			expectedError: "\nInvalid Command: add requires source name" + UsageMessage,
		},
		"Verify parses other arguments passed along with `source add`": {
			input:          []string{"source", "add", "source-name"},
			expectedOutput: &models.Command{Primary: "source", Secondary: "add", Arguments: []string{"source-name"}},
		},
		"Verify throws error when no source name is provided with `source remove`": {
			isNegative:    true,
			input:         []string{"source", "remove"},
			expectedError: "\nInvalid Command: remove requires source name" + UsageMessage,
		},
		"Verify parses other arguments passed along with `source remove`": {
			input:          []string{"source", "remove", "source-name"},
			expectedOutput: &models.Command{Primary: "source", Secondary: "remove", Arguments: []string{"source-name"}},
		},
		"Verify parses other arguments passed along with `source list`": {
			input:          []string{"source", "list"},
			expectedOutput: &models.Command{Primary: "source", Secondary: "list"},
		},
	}

	runTableTest(t, tests)
}

func TestHelpCommand(t *testing.T) {
	_, err := Parse([]string{"restoros", "help"})
	assert.EqualError(t, err, UsageMessage)
}

func runTableTest(t *testing.T, tests map[string]Test) {
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			input := append([]string{"restoros"}, test.input...)
			actualOutput, actualError := Parse(input)
			if !test.isNegative {
				assert.Equal(t, test.expectedOutput, actualOutput)
			} else {
				assert.EqualError(t, actualError, test.expectedError)
			}
		})
	}
}
