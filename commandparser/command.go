package commandparser

type Command struct {
	handler func([]string) bool
	args []string
}

func (command Command) Exec() bool {
	return command.handler(command.args);
}