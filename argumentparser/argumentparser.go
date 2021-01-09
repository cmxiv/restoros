package argumentparser

import (
	"errors"
	"fmt"
)

type Command struct {
	primary   string
	secondary string
	arguments []string
}

type parserFunc = func([]string) (*Command, error)

var commandMap = map[string]parserFunc{
	"install": parsePrimaryWithArguments("install"),
	"update":  parsePrimaryWithArguments("update"),
	"remove":  parsePrimaryWithArguments("remove"),
	"purge":   parsePrimaryWithArguments("purge"),
	"restore": parsePrimaryOnly("restore"),
	"reset":   parsePrimaryOnly("reset"),
	"list":    parsePrimaryOnly("list"),
	"source":  parsePrimaryWithSecondaryAndArgs("source", sourceSecondaryCommandMap),
}

var sourceSecondaryCommandMap = map[string]parserFunc{
	"list":   parseSecondaryOnly("list"),
	"add":    parseSecondaryWithArguments("add"),
	"remove": parseSecondaryWithArguments("remove"),
}

func Parse(arguments []string) (*Command, error) {

	if len(arguments) < 2 {
		return nil, errorMessage(nil, nil)
	}

	primary := arguments[1]
	if primary == "help" {
		return nil, errorMessage(nil, nil)
	}

	if commandMap[primary] == nil {
		return nil, errorMessage(&primary, nil)
	}
	return commandMap[primary](arguments[2:])
}

func parsePrimaryOnly(primary string) parserFunc {
	return func(args []string) (*Command, error) {
		return &Command{primary: primary}, nil
	}
}

func parseSecondaryOnly(secondary string) parserFunc {
	return func(arguments []string) (*Command, error) {
		return &Command{secondary: secondary}, nil
	}
}

func parsePrimaryWithArguments(primary string) parserFunc {
	return func(arguments []string) (*Command, error) {
		if len(arguments) < 1 {
			err := "requires package name"
			return nil, errorMessage(&primary, &err)
		}
		return &Command{
			primary:   primary,
			arguments: arguments,
		}, nil
	}
}

func parseSecondaryWithArguments(secondary string) parserFunc {
	return func(arguments []string) (*Command, error) {
		if len(arguments) < 1 {
			err := "requires source name"
			return nil, errorMessage(&secondary, &err)
		}
		return &Command{
			secondary: secondary,
			arguments: arguments,
		}, nil
	}
}

func parsePrimaryWithSecondaryAndArgs(primary string, secondaryMap map[string]parserFunc) parserFunc {
	return func(arguments []string) (*Command, error) {

		if len(arguments) < 1 {
			err := "requires a sub-option"
			return nil, errorMessage(&primary, &err)
		}

		secondary := arguments[0]
		if secondaryMap[secondary] == nil {
			return nil, errorMessage(&primary, &secondary)
		}

		subCommand, err := secondaryMap[secondary](arguments[1:])
		if err != nil {
			return nil, err
		}

		return &Command{
			primary:   primary,
			secondary: subCommand.secondary,
			arguments: subCommand.arguments,
		}, nil
	}
}

func errorMessage(primary *string, secondary *string) error {

	if primary == nil {
		return errors.New(UsageMessage)
	}

	secCommand := ""
	if secondary != nil {
		secCommand = *secondary
	}
	invalidCommand := fmt.Sprintf("\nInvalid Command: %s %s", *primary, secCommand)

	return errors.New(invalidCommand + UsageMessage)
}

const UsageMessage = `

Usage: restoros [options] <sub-options> <package>

Where options include:

	install: Installs the provided package by searching for them in the configured sources.

	update: Updates the provided package if any available

	remove: Removes the provided package from the system.

	purge: Removes the provided package from the system and purges any records from Restoros as well

	restore: Restores the state of the system to the restoros configurations

	reset: Restores the state of the system to when restoros was first installed

	source: This option let's user manage sources. The following are the sub-commands:
		add: Add a source to restoros configuration
		remove: Remove a source from restoros configuration
		list: List all the present sources. The search for the packages will be done in the listed order
		reorder: Reorder the currently configured sources

	list: Lists all the present packages managed by restoros
`
