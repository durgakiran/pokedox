package pokemon

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"durgakiran.dev/pokedox/cache"
)

var CatchUrl = Url{
	Endpoint:    "pokemon/%s",
	Type:        Methods["get"],
	Description: "Get details about pokemon",
}

var pokemonCache = cache.NewCache(5 * time.Minute)

func GetPokemon(pokemonName string) Pokemon {
	url := fmt.Sprintf(BaseUrl+CatchUrl.Endpoint, pokemonName)
	pokeMon := Pokemon{}

	cache, ok := pokemonCache.Get(url)

	if ok {
		// fmt.Printf("Getting location from cache with url: %s", url)
		err := json.Unmarshal(cache, &pokeMon)
		if err != nil {
			log.Fatal(err)
		}
		return pokeMon
	}
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(body, &pokeMon)
	if err != nil {
		log.Fatal(err)
	}

	pokemonCache.Add(url, body)

	return pokeMon
}

var pokemonCaught = make(map[string]Pokemon)

func SaveCaughtPokemon(pokemonName string) {
	pokemonCaught[pokemonName] = GetPokemon(pokemonName)
}

func IsPokemonCaught(pokemonName string) bool {
	_, ok := pokemonCaught[pokemonName]
	return ok
}

func GetCaughtPokemon(pokemonName string) Pokemon {
	return pokemonCaught[pokemonName]
}

func GetAllCaughtPokemon() map[string]Pokemon {
	return pokemonCaught
}
