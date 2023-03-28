package commands

import (
	"fmt"

	"durgakiran.dev/pokedox/core"
	"durgakiran.dev/pokedox/pokemon"
)

var commandLocation = core.CliCommnad{
	Name:        "map",
	Description: "Get the locations of the pokemon",
	Usage:       "map",
	Command:     Location,
}

var commandLocationPrev = core.CliCommnad{
	Name:        "mapb",
	Description: "Get the locations of the pokemon",
	Usage:       "mapb",
	Command:     LocationPrev,
}

func DeclareLocation() {
	core.CommandConfig.Commands["map"] = commandLocation
}

func DeclareLocationPrev() {
	core.CommandConfig.Commands["mapb"] = commandLocationPrev
}

func Location(args ...string) string {
	data := pokemon.GetLocation()

	for _, v := range data.Results {
		fmt.Println(v.Name)
	}

	return ""
}

func LocationPrev(args ...string) string {
	data, err := pokemon.PreviousLocation()

	if err != nil {
		fmt.Println(err)
		return ""
	}

	for _, v := range data.Results {
		fmt.Println(v.Name)
	}

	return ""
}
