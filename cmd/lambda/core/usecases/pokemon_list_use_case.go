package usecases

import (
	"encoding/json"
	"log"
	"testProject/core/entities"
	"testProject/utils"
)

func GetPokemonList() (entities.PokemonList, error) {

	url := "https://pokeapi.co/api/v2/pokemon?limit=151"
	var pokemonList entities.PokemonList

	resp, err := utils.ApiRequest(url)

	err = json.NewDecoder(resp).Decode(&pokemonList)

	log.Println(pokemonList)

	return pokemonList, err
}
