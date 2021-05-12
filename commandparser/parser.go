package commandparser

func Parse(args []string) (*Command, bool) {
	return createCommandTree().parse(args)
}

func createCommandTree() *node {
	return &node{
		children: []*node{
			{
				argument: "import",
				children: []*node{
					{command: &Command{handler: falseFunc}},
				},
			},
			{
				argument: "install",
				children: []*node{
					{command: &Command{handler: falseFunc}},
				},
			},
			{
				argument: "update",
				children: []*node{
					{command: &Command{handler: falseFunc}},
				},
			},
			{
				argument: "remove",
				children: []*node{
					{command: &Command{handler: falseFunc}},
				},
			},
			{
				argument: "purge",
				children: []*node{
					{command: &Command{handler: falseFunc}},
				},
			},
			{
				argument: "restore",
				children: []*node{
					{command: &Command{handler: falseFunc}},
				},
			},
			{
				argument: "reset",
				children: []*node{
					{command: &Command{handler: falseFunc}},
				},
			},
			{
				argument: "source",
				children: []*node{
					{
						argument: "add",
						children: []*node{
							{command: &Command{handler: falseFunc}},
						},
					},
					{
						argument: "remove",
						children: []*node{
							{command: &Command{handler: falseFunc}},
						},
					},
					{
						argument: "list",
						children: []*node{
							{command: &Command{handler: falseFunc}},
						},
					},
				},
			},
			{
				argument: "config",
				children: []*node{
					{
						argument: "init",
						children: []*node{
							{command: &Command{handler: falseFunc}},
						},
					},
					{
						argument: "sync",
						children: []*node{
							{command: &Command{handler: falseFunc}},
						},
					},
					{
						argument: "origin",
						children: []*node{
							{command: &Command{handler: falseFunc}},
						},
					},
				},
			},
			{
				argument: "list",
				children: []*node{
					{command: &Command{handler: falseFunc}},
				},
			},
		},
	}
}

func falseFunc(args []string) bool {
	return false
}

type node struct {
	argument string
	children []*node
	command  *Command
}

func (n *node) parse(args []string) (*Command, bool) {
	tmp := n
	var command *Command = nil
	commandIndex := 0
	for _, arg := range args {
		for _, child := range tmp.children {
			if arg == child.argument {
				commandIndex++
				tmp = child
				break
			}
		}
		if len(tmp.children) == 1 && tmp.children[0].command != nil {
			command = &Command{
				handler: tmp.children[0].command.handler,
				args: args[commandIndex:],
			}
			break
		}
	}
	return command, command != nil
}
