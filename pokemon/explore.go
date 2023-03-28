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

var ExploreUrl = Url{
	Endpoint:    "location-area/%s",
	Type:        Methods["get"],
	Description: "Location areas are sections of areas, such as floors in a building or cave. Each area has its own set of possible Pokémon encounters.",
}

var locationExploreCache = cache.NewCache(5 * time.Minute)

func ExploreArea(location string) LocationArea {
	url := fmt.Sprintf(BaseUrl+ExploreUrl.Endpoint, location)
	locationArea := LocationArea{}

	cache, ok := locationExploreCache.Get(url)

	if ok {
		// fmt.Printf("Getting location from cache with url: %s", url)
		err := json.Unmarshal(cache, &locationArea)
		if err != nil {
			log.Fatal(err)
		}
		return locationArea
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
	err = json.Unmarshal(body, &locationArea)
	if err != nil {
		log.Fatal(err)
	}

	locationExploreCache.Add(url, body)

	return locationArea
}
