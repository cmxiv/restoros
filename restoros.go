package main

import (
	"fmt"
	"os"
	"restoros/argumentparser"
	"restoros/config"
	"restoros/handler"
	"restoros/sourcehandler"
)

func main() {

	cfg, err := config.Read()
	if err != nil {
		printAndExit(err)
	}

	command, err := argumentparser.Parse(os.Args)
	if err != nil {
		printAndExit(err)
	}

	var hndlr handler.Handler
	if hndlr, err = handler.GetHandler(command); err != nil {
		printAndExit(err)
	}

	cfg = hndlr.Handle(cfg)
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
