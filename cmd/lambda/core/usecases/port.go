package usecases

import (
	"context"
	"testProject/core/entities"
)

type IPokemonRepo interface {
	GetPokemon(ctx context.Context, pkm entities.Pokemon) (entities.Pokemon, error)
}

type IPokemonAttrsRepo interface {
	SaveItem(iapPurchaseId int, payload map[string]interface{}) error
}
