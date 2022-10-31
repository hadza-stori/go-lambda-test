package usecases

import (
	"context"
	"encoding/json"
	"errors"
	"testProject/core/entities"
	"testProject/utils"
)

type PokemonUseCase struct {
	PokemonRepo IPokemonRepo
}

func GetPokemonUseCase(pokemonRepo IPokemonRepo) PokemonUseCase {
	return PokemonUseCase{
		PokemonRepo: pokemonRepo,
	}
}

func (uc *PokemonUseCase) Execute(ctx context.Context, pkm entities.Pokemon) (entities.Pokemon, error) {
	url := "https://pokeapi.co/api/v2/pokemon/" + pkm.Name
	var pokemon entities.Pokemon

	resp, _ := utils.ApiRequest(url)

	if resp == nil {
		return pokemon, errors.New("pokemon not found")
	}

	json.NewDecoder(resp).Decode(&pokemon)

	return pokemon, nil
}
