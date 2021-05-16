package commandparser

import "restoros/handler"

type Command struct {
	handler handler.Handler
	args []string
}

func (command Command) Exec() error {
	return command.handler.Handle(command.args)
}