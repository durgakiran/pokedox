package commands

import (
	"fmt"

	"durgakiran.dev/pokedox/core"
	"durgakiran.dev/pokedox/pokemon"
)

var commandExplore = core.CliCommnad{
	Name:        "explore",
	Description: "explore location in the map",
	Usage:       "explore <location>",
	Command:     Explore,
}

func DeclareExplore() {
	core.CommandConfig.Commands["explore"] = commandExplore
}

func Explore(args ...string) string {
	location := args[0]
	fmt.Println("Exploring Area", location, "...")
	data := pokemon.ExploreArea(location)

	fmt.Println("Pokemon found:")
	for _, v := range data.PokemonEncounters {
		fmt.Println("-", v.Pokemon.Name)
	}

	return ""
}
