package pokemon

type NamedAPIResource struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Name struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type GenerationGameIndex struct {
	GameIndex  int              `json:"game_index"`
	Generation NamedAPIResource `json:"generation"`
}

type Location struct {
	Id          int                   `json:"id"`
	Name        string                `json:"name"`
	Region      NamedAPIResource      `json:"region"`
	Names       []Name                `json:"names"`
	GameIndices []GenerationGameIndex `json:"game_indices"`
	Areas       []NamedAPIResource    `json:"areas"`
}

type Response[T any] struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []T    `json:"results"`
}

type EncounterVersionDetail struct {
	Rate    int              `json:"rate"`
	Version NamedAPIResource `json:"version"`
}

type VersionEncounterDetail struct {
	Text         string           `json:"text"`
	Language     NamedAPIResource `json:"language"`
	VersionGroup NamedAPIResource `json:"version_group"`
}

type PokemonEncounter struct {
	Pokemon        NamedAPIResource         `json:"pokemon"`
	VersionDetails []VersionEncounterDetail `json:"version_details"`
}

type EncounterMethodRate struct {
	EncounterMethod NamedAPIResource         `json:"encounter_method"`
	VersionDetails  []EncounterVersionDetail `json:"version_details"`
}

type LocationArea struct {
	Id                   int                   `json:"id"`
	Name                 string                `json:"name"`
	GameIndex            int                   `json:"game_index"`
	EncounterMethodRates []EncounterMethodRate `json:"encounter_method_rates"`
	Location             NamedAPIResource      `json:"location"`
	Names                []Name                `json:"names"`
	PokemonEncounters    []PokemonEncounter    `json:"pokemon_encounters"`
}

type PokeMonAbility struct {
	IsHidden bool             `json:"is_hidden"`
	Slot     int              `json:"slot"`
	Ability  NamedAPIResource `json:"ability"`
}

type PokemonHeldItemVersion struct {
	Version NamedAPIResource `json:"version"`
	Rarity  int              `json:"rarity"`
}

type PokemonHeldItem struct {
	Name           string           `json:"name"`
	Item           NamedAPIResource `json:"item"`
	VersionDetails []interface{}    `json:"version_details"`
}

type PokemonMoveVersion struct {
	MoveLearnMethod NamedAPIResource `json:"move_learn_method"`
	VersionGroup    NamedAPIResource `json:"version_group"`
	LevelLearnedAt  int              `json:"level_learned_at"`
}

type PokemonMove struct {
	Move                NamedAPIResource     `json:"move"`
	VersionGroupDetails []PokemonMoveVersion `json:"version_group_details"`
}

type PokemonSprites struct {
	FrontDefault     string `json:"front_default"`
	FrontShiny       string `json:"front_shiny"`
	FrontFemale      string `json:"front_female"`
	FrontShinyFemale string `json:"front_shiny_female"`
	BackDefault      string `json:"back_default"`
	BackShiny        string `json:"back_shiny"`
	BackFemale       string `json:"back_female"`
	BackShinyFemale  string `json:"back_shiny_female"`
}

type PokemonStat struct {
	Stat     NamedAPIResource `json:"stat"`
	Effort   int              `json:"effort"`
	BaseStat int              `json:"base_stat"`
}

type PokemonType struct {
	Slot int              `json:"slot"`
	Type NamedAPIResource `json:"type"`
}

type Pokemon struct {
	Id                     int                   `json:"id"`
	Name                   string                `json:"name"`
	BaseExperience         int                   `json:"base_experience"`
	Height                 int                   `json:"height"`
	IsDefault              bool                  `json:"is_default"`
	Order                  int                   `json:"order"`
	Weight                 int                   `json:"weight"`
	Abilities              []PokeMonAbility      `json:"abilities"`
	Forms                  []NamedAPIResource    `json:"forms"`
	GameIndices            []GenerationGameIndex `json:"game_indices"`
	HeldItems              []PokemonHeldItem     `json:"held_items"`
	LocationAreaEncounters string                `json:"location_area_encounters"`
	Moves                  []PokemonMove         `json:"moves"`
	PastTypes              []NamedAPIResource    `json:"past_types"`
	Sprites                PokemonSprites        `json:"sprites"`
	Stats                  []PokemonStat         `json:"stats"`
	Types                  []PokemonType         `json:"types"`
}
