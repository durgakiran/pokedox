package pokemon

var BaseUrl = "https://pokeapi.co/api/v2/"

var Methods = map[string]string{
	"get":  "GET",
	"post": "POST",
}

type Url struct {
	Endpoint    string
	Type        string
	Description string
}
