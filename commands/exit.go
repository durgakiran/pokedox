package commands

import (
	"os"

	"durgakiran.dev/pokedox/core"
)

var commandExit = core.CliCommnad{
	Name:        "exit",
	Description: "exit from the program",
	Usage:       "exit",
	Command:     Exit,
}

func DeclareExit() {
	core.CommandConfig.Commands["exit"] = commandExit
}

func Exit(args ...string) string {
	os.Exit(0)
	return ""
}
