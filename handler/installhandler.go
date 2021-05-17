package handler

import (
	"flag"
	"fmt"
	"os"
	"restoros/configurationmanager"
	"restoros/sourcemanager"
	"strings"
)

type InstallHandler struct {
	SourceManager sourcemanager.SourceManager
	Manager       configurationmanager.Manager
	RepoManager   configurationmanager.RepositoryManager
}

func (handler *InstallHandler) Handle(args []string) error {
	parsedArguments, err := handler.parseArguments(args)
	if err != nil {
		return fmt.Errorf(INSTALL_USAGE)
	}

	source, err := handler.findSource(parsedArguments.source, parsedArguments.packageName)
	if err != nil {
		return err
	}

	if err = source.Install(); err != nil {
		return fmt.Errorf("unable to install %s from source %s", parsedArguments.packageName, parsedArguments.source)
	}

	pkg := configurationmanager.Package{
		Source:  source.Name(),
		Version: parsedArguments.version,
		Name:    parsedArguments.packageName,
	}
	if err = handler.Manager.AddPackage(pkg); err != nil {
		return fmt.Errorf("unable to add package to configuration")
	}

	return handler.RepoManager.Sync()
}

type parsedArguments struct {
	source      string
	version     string
	packageName string
}

func (handler *InstallHandler) parseArguments(args []string) (parsedArguments, error) {
	r, w, _ := os.Pipe()
	defer func(read *os.File, write *os.File) {
		read.Close()
		write.Close()
	}(r, w)

	var (
		version string
		source  string
	)

	flagSet := flag.NewFlagSet("", flag.ContinueOnError)
	flagSet.SetOutput(w)
	flagSet.StringVar(&source, "s", "", "")
	flagSet.StringVar(&source, "source", "", "")
	flagSet.StringVar(&version, "v", "default", "")
	flagSet.StringVar(&version, "version", "default", "")
	err := flagSet.Parse(args)

	return parsedArguments{
		source:      source,
		version:     version,
		packageName: flagSet.Args()[0],
	}, err
}

func (handler *InstallHandler) findSource(source string, packageName string) (sourcemanager.Source, error) {
	if source != "" {
		return handler.SourceManager.FindByName(source), nil
	}

	sources := handler.SourceManager.FindByPackage(packageName)
	if len(sources) == 1 {
		return sources[0], nil
	}

	if len(sources) < 1 {
		return nil, fmt.Errorf("unable to find package in any configured source")
	}

	sourceNames := []string{}
	for _, source := range sources {
		sourceNames = append(sourceNames, source.Name())
	}
	return nil, fmt.Errorf(
		"found multiple sources (see below) with package %s, use --source to select your choice.\n%s",
		packageName,
		strings.Join(sourceNames, "\n"),
	)
}

const INSTALL_USAGE = `
Available flags:
	
	-s, --source - Set the source for the installation

	-v, --version - Set the version which needs to be installed`
