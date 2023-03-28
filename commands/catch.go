package commands

import (
	"fmt"
	"math/rand"

	"durgakiran.dev/pokedox/core"
	"durgakiran.dev/pokedox/pokemon"
)

var commandCatch = core.CliCommnad{
	Name:        "catch",
	Description: "catch pokemon",
	Usage:       "catch <pokemon name>",
	Command:     Catch,
}

func DeclareCatch() {
	core.CommandConfig.Commands["catch"] = commandCatch
}

var effortBase = 10

func Catch(args ...string) string {
	pokemonName := args[0]
	caught := pokemon.IsPokemonCaught(pokemonName)
	if caught {
		fmt.Println(pokemonName, "is already caught!")
		return ""
	}

	fmt.Println("Throwing a Pokeball at", pokemonName, "...")
	data := pokemon.GetPokemon(pokemonName)

	difficutly := data.BaseExperience

	energy := rand.Intn(effortBase)

	if energy > difficutly {
		fmt.Println(pokemonName, "caught!")
		pokemon.SaveCaughtPokemon(pokemonName)
	} else {
		fmt.Println(pokemonName, "escaped!")
		effortBase *= 2
	}

	return ""
}
