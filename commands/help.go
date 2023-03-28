package commands

import (
	"fmt"

	"durgakiran.dev/pokedox/core"
)

var command = core.CliCommnad{
	Name:        "help",
	Description: "List all the commands",
	Usage:       "help",
	Command:     Help,
}

func DeclareHelp() {
	core.CommandConfig.Commands["help"] = command
}

func Help(args ...string) string {
	fmt.Println("Welocom to pokedox!")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println()
	for k, v := range core.CommandConfig.Commands {
		fmt.Println(k, ":", v.Description)
	}
	return ""
}
