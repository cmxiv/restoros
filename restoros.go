package main

import (
	"fmt"
	"os"
	"restoros/argumentparser"
)

func main() {

	command, err := argumentparser.Parse(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	fmt.Println(command)

}
