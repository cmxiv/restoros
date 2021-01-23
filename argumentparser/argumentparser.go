package argumentparser

import (
	"errors"
	"fmt"
	"restoros/models"
)

type parserFunc = func([]string) (*models.Command, error)

var commandMap = map[string]parserFunc{
	"install": parsePrimaryWithArguments("install"),
	"update":  parsePrimaryWithArguments("update"),
	"remove":  parsePrimaryWithArguments("remove"),
	"purge":   parsePrimaryWithArguments("purge"),
	"restore": parsePrimaryOnly("restore"),
	"reset":   parsePrimaryOnly("reset"),
	"list":    parsePrimaryOnly("list"),
	"source":  parsePrimaryWithSecondaryAndArgs("source", sourceSecondaryCommandMap),
	"config":  parsePrimaryWithSecondaryAndArgs("config", configSecondaryCommandMap),
}

var sourceSecondaryCommandMap = map[string]parserFunc{
	"list":   parseSecondaryOnly("list"),
	"add":    parseSecondaryWithArguments("add", "requires source name"),
	"remove": parseSecondaryWithArguments("remove", "requires source name"),
}

var configSecondaryCommandMap = map[string]parserFunc{
	"init":   parseSecondaryOnly("init"),
	"sync":   parseSecondaryOnly("sync"),
	"origin": parseSecondaryWithArguments("origin", ""),
}

func Parse(arguments []string) (*models.Command, error) {

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
	return func(args []string) (*models.Command, error) {
		return &models.Command{Primary: primary}, nil
	}
}

func parseSecondaryOnly(secondary string) parserFunc {
	return func(arguments []string) (*models.Command, error) {
		return &models.Command{Secondary: secondary}, nil
	}
}

func parsePrimaryWithArguments(primary string) parserFunc {
	return func(arguments []string) (*models.Command, error) {
		if len(arguments) < 1 {
			err := "requires package name"
			return nil, errorMessage(&primary, &err)
		}
		return &models.Command{
			Primary:   primary,
			Arguments: arguments,
		}, nil
	}
}

func parseSecondaryWithArguments(secondary string, errMessage string) parserFunc {
	return func(arguments []string) (*models.Command, error) {
		if len(arguments) < 1 && errMessage != "" {
			return nil, errorMessage(&secondary, &errMessage)
		}
		return &models.Command{
			Secondary: secondary,
			Arguments: arguments,
		}, nil
	}
}

func parsePrimaryWithSecondaryAndArgs(primary string, secondaryMap map[string]parserFunc) parserFunc {
	return func(arguments []string) (*models.Command, error) {

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

		return &models.Command{
			Primary:   primary,
			Secondary: subCommand.Secondary,
			Arguments: subCommand.Arguments,
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

	source: This option let's user manage sources. Following are the sub-commands:
		add: Add a source to restoros configuration
		remove: Remove a source from restoros configuration
		list: List all the present sources. The search for the packages will be done in the listed order
		reorder: Reorder the currently configured sources
	
	config: This option let's user to manage the installation's configurations. Following are the sub-commands:
		init: Initializes configurations for restoros
		sync: Syncronizes with the origin all the changes in the configuration
		origin: Set the origin to the provided github repository; Returns the current origin if not provided with an argument

	list: Lists all the present packages managed by restoros
`
