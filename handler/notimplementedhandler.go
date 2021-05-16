package handler

import "fmt"

type NotImplementedHandler struct{}

func (handler *NotImplementedHandler) Handle(args []string) error {
	return fmt.Errorf("command not implemented, yet!")
}
