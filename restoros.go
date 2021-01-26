package main

import (
	"fmt"
	"os"
	"restoros/argumentparser"
	"restoros/config"
	"restoros/handler"
)

func main() {

	command, err := argumentparser.Parse(os.Args)
	if err != nil {
		printAndExit(err)
	}

	cfg, err := config.Read(command)
	if err != nil {
		printAndExit(err)
	}

	var hndlr handler.Handler
	if hndlr, err = handler.GetHandler(command); err != nil {
		printAndExit(err)
	}

	cfg, err = hndlr.Handle(cfg)
	if cfg == nil {
		printAndExit(fmt.Errorf("Nil configuration"))
	}

	if err != nil {
		printAndExit(err)
	}

	if cfg.Modified {
		if err = config.Write(cfg); err != nil {
			printAndExit(err)
		}
	}

}

func printAndExit(err error) {
	fmt.Println(err)
	os.Exit(-1)
}
