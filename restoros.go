package main

import (
	"fmt"
	"os"
	"restoros/commandparser"
)

func main() {
	command, valid := commandparser.Parse(os.Args[1:])
	if !valid {
		fmt.Print(UsageMessage)
		os.Exit(-1)
	}
	if err := command.Exec(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

const UsageMessage = `
Usage: restoros [options] <sub-options> <package>

	Where options include:

		import: Import the existing installed packages

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
