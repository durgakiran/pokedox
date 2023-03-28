package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"durgakiran.dev/pokedox/commands"
	"durgakiran.dev/pokedox/core"
)

func init() {
	// Declare all the commands here
	commands.DeclareHelp()
	commands.DeclareExit()
	commands.DeclareLocation()
	commands.DeclareLocationPrev()
	commands.DeclareExplore()
	commands.DeclareCatch()
	commands.DeclareInspect()
	commands.DeclarePokedox()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("pokedox > ")
		if !scanner.Scan() {
			break
		}

		line := scanner.Text()
		args := strings.Split(line, " ")
		eval(args[0], args[1:]...)
	}
}

func eval(command string, args ...string) {
	val, ok := core.CommandConfig.Commands[command]

	if !ok {
		fmt.Println("Command not found")
		return
	}

	fmt.Println(val.Command(args...))
}
