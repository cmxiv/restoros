package commandparser

import (
	"restoros/configurationmanager"
	"restoros/handler"
	"restoros/handler/confighandler"
)

func Parse(args []string) (*Command, bool) {
	return createCommandTree().parse(args)
}

func createCommandTree() *node {
	manager := &configurationmanager.Manager{}
	repositoryManager := &configurationmanager.RepositoryManager{Path: configurationmanager.RestorosDirectory()}
	notImplementedHandler := &handler.NotImplementedHandler{}
	return &node{
		children: []*node{
			{
				argument: "import",
				children: []*node{
					{command: &Command{handler: notImplementedHandler}},
				},
			},
			{
				argument: "install",
				children: []*node{
					{command: &Command{handler: notImplementedHandler}},
				},
			},
			{
				argument: "update",
				children: []*node{
					{command: &Command{handler: notImplementedHandler}},
				},
			},
			{
				argument: "remove",
				children: []*node{
					{command: &Command{handler: notImplementedHandler}},
				},
			},
			{
				argument: "purge",
				children: []*node{
					{command: &Command{handler: notImplementedHandler}},
				},
			},
			{
				argument: "restore",
				children: []*node{
					{command: &Command{handler: notImplementedHandler}},
				},
			},
			{
				argument: "reset",
				children: []*node{
					{command: &Command{handler: notImplementedHandler}},
				},
			},
			{
				argument: "source",
				children: []*node{
					{
						argument: "add",
						children: []*node{
							{command: &Command{handler: notImplementedHandler}},
						},
					},
					{
						argument: "remove",
						children: []*node{
							{command: &Command{handler: notImplementedHandler}},
						},
					},
					{
						argument: "list",
						children: []*node{
							{command: &Command{handler: notImplementedHandler}},
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
							{
								command: &Command{
									handler: &confighandler.ConfigInitHandler{
										Manager:     manager,
										RepoManager: repositoryManager,
									},
								},
							},
						},
					},
					{
						argument: "sync",
						children: []*node{
							{
								command: &Command{
									handler: &confighandler.ConfigSyncHandler{
										RepoManager: repositoryManager,
									},
								},
							},
						},
					},
					{
						argument: "origin",
						children: []*node{
							{
								command: &Command{
									handler: &confighandler.ConfigOriginHandler{
										RepoManager: repositoryManager,
									},
								},
							},
						},
					},
				},
			},
			{
				argument: "list",
				children: []*node{
					{command: &Command{handler: notImplementedHandler}},
				},
			},
		},
	}
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
				args:    args[commandIndex:],
			}
			break
		}
	}
	return command, command != nil
}
