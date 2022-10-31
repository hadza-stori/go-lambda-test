package mocks

import (
	"context"
	"github.com/stretchr/testify/mock"
	"testProject/core/entities"
)

type PokemonTestMock struct {
	mock.Mock
}

func (m *PokemonTestMock) GetPokemon(ctx context.Context, pokemon entities.Pokemon) (entities.Pokemon, error) {
	args := m.Called(ctx, pokemon)
	return args.Get(0).(entities.Pokemon), args.Error(1)
}
