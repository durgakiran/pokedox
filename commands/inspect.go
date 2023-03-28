package commands

import (
	"fmt"

	"durgakiran.dev/pokedox/core"
	"durgakiran.dev/pokedox/pokemon"
)

var commandInspect = core.CliCommnad{
	Name:        "inspect",
	Description: "inspect pokemon if already caught",
	Usage:       "inspect <pokemon name>",
	Command:     Inspect,
}

func DeclareInspect() {
	core.CommandConfig.Commands["inspect"] = commandInspect
}

func Inspect(args ...string) string {
	pokemonName := args[0]
	caught := pokemon.IsPokemonCaught(pokemonName)
	if !caught {
		fmt.Println("you have not caught that pokemon yet!")
		return ""
	}

	pokemonDesc := pokemon.GetPokemon(pokemonName)

	fmt.Println("Name:", pokemonDesc.Name)
	fmt.Println("Height:", pokemonDesc.Height)
	fmt.Println("Weight:", pokemonDesc.Weight)
	fmt.Println("Stats:")
	for _, v := range pokemonDesc.Stats {
		fmt.Println("	-", v.Stat.Name+":", v.BaseStat)
	}
	fmt.Println("Types:")
	for _, v := range pokemonDesc.Types {
		fmt.Println("	-", v.Type.Name)
	}

	return ""
}
