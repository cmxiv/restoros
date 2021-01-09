package main

import (
	"fmt"
	"restoros/config"
)

func main() {

	json, err := config.Read()
	fmt.Printf("%+v %v\n", json, err)

}
