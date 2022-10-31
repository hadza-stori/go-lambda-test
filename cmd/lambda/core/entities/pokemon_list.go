package entities

type PokemonList struct {
	Count    int               `json:"count"`
	Next     string            `json:"next"`
	Previous string            `json:"previous"`
	Results  []PokemonListItem `json:"results"`
}
