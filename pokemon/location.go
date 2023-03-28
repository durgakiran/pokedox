package pokemon

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"time"

	"durgakiran.dev/pokedox/cache"
)

var LocationUrl = Url{
	Endpoint:    "location-area/?offset=0&limit=20",
	Type:        Methods["get"],
	Description: "Locations that can be visited within the games. Locations make up sizable portions of regions, like cities or routes.",
}

var nextLocations = ""
var previousLocations = ""

var locationCache = cache.NewCache(5 * time.Minute)

func GetLocation() Response[LocationArea] {

	url := BaseUrl + LocationUrl.Endpoint

	if nextLocations != "" {
		url = nextLocations
	}

	cache, ok := locationCache.Get(url)

	location := Response[LocationArea]{}

	if ok {
		// fmt.Printf("Getting location from cache with url: %s", url)
		err := json.Unmarshal(cache, &location)
		if err != nil {
			log.Fatal(err)
		}
	} else {
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
		err = json.Unmarshal(body, &location)
		if err != nil {
			log.Fatal(err)
		}
		locationCache.Add(url, body)
	}

	nextLocations = location.Next
	previousLocations = location.Previous

	return location
}

func PreviousLocation() (Response[LocationArea], error) {

	url := ""

	if previousLocations != "" {
		url = previousLocations
	} else {
		return Response[LocationArea]{}, errors.New("no previous locations")
	}

	cache, ok := locationCache.Get(url)

	location := Response[LocationArea]{}

	if ok {
		// fmt.Printf("Getting location from cache with url: %s", url)
		err := json.Unmarshal(cache, &location)
		if err != nil {
			log.Fatal(err)
		}
	} else {

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

		err = json.Unmarshal(body, &location)

		if err != nil {
			log.Fatal(err)
		}
		locationCache.Add(url, body)
	}

	nextLocations = location.Next
	previousLocations = location.Previous

	return location, nil
}
