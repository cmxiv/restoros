# Restoros

![Restoros](https://github.com/cmxiv/restoros/workflows/restoros/badge.svg)

## What is it?

Restoros is a package manager aggregator. You can install system and development packages using it and sync the installations to a GitHub repository making it easy for restoring you development setup with one single command.

## How does it work?

Restoros works as a go-between multiple package managers and maintain a record for the installations in its configuration file. This configuration file is being sync to a remote repository of the user's choice (currently only support GitHub).

If the user is settings up a new system or recovering an existing system, they would have just run
```
$ restoros config origin <path to config repository>
$ resotors restore
```

## Usage

```
restoros [options] <sub-options> <package>
```

| Options | Description |
|---------|-------------|
| Import | Import all the possible packages from the current system to restoros |
| Install | Install package by searching for them in the configured sources. Return with options if multiple hits in sources or fails if nothing found |
| update | Updates package if any available |
| remove | Removes package if installed (doesn't remove from configuration) |
| purge | Removes package and purges any records from restoros configuration |
| restore | Restores packages using the remote configuration (see config option for more details) |
| reset | Restors the state of the system to when restoros was first installed (under scrutiny, might not make it in the releases) |
| source | This option is used to manage the sources (the package managers) from where the packages are being installed/managed. See the table below for the sub-options for source|
| config | This option is used to manage restoros's configurations. See the table below for the sub-options for config |
| list | Lists all the installed packages being managed by restoros |

### Source sub-options

Following are the source sub-options:
| Sub-options | Description |
|-----------|-------------|
| add | Add a source to restoros configuration |
| remove | Remove a source from restoros configuration|
| list | List all the present sources. The search for the pakcages will be done in the listed order |

### Config sub-options

Following are the config sub-options:
| Sub-options | Description |
|-----------|-------------|
|init|Initializes configuraitons for restoros|
|sync|Syncronizes with the origin all the changes in the configuration|
|origin|Set the origin to the provided github repository; Returns the current origin if not provided with an argument|


## Contributions

Thanks for choosing my project for your contribution. Following are somethings to consider when creating a PR:

- It's prefered that you follow TDD, writing tests supporting the code that you are about to submit
- Raise PR with proper title and message along with comments in PR whereever you feel is necessary

Before download and setup of the project, make sure that you have Go installed. Minimum requirement is Go version 1.15. To download and start development:
```
git clone git@github.com:cmxiv/restoros.git
cd restoros
make test
```