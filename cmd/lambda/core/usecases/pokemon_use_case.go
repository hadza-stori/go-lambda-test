package usecases

import (
	"encoding/json"
	"errors"
	"testProject/core/entities"
	"testProject/utils"
)

func GetPokemon(pkm entities.Pokemon) (entities.Pokemon, error) {

	url := "https://pokeapi.co/api/v2/pokemon/" + pkm.Name
	var pokemon entities.Pokemon

	resp, _ := utils.ApiRequest(url)

	if resp == nil {
		return pokemon, errors.New("pokemon not found")
	}

	json.NewDecoder(resp).Decode(&pokemon)

	return pokemon, nil
}
