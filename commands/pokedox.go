package commands

import (
	"fmt"

	"durgakiran.dev/pokedox/core"
	"durgakiran.dev/pokedox/pokemon"
)

var commandPokedox = core.CliCommnad{
	Name:        "Pokedox",
	Description: "Pokedox pokemon if already caught",
	Usage:       "pokedox <pokemon name>",
	Command:     Pokedox,
}

func DeclarePokedox() {
	core.CommandConfig.Commands["pokedox"] = commandPokedox
}

func Pokedox(args ...string) string {

	pokemonDesc := pokemon.GetAllCaughtPokemon()
	fmt.Println("Your Pokedox:")
	for _, v := range pokemonDesc {
		fmt.Println("-", v.Name)
	}

	return ""
}
