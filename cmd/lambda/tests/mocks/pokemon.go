package mocks

import (
	"github.com/stretchr/testify/mock"
	"testProject/core/entities"
)

type PokemonTestMock struct {
	mock.Mock
}

func (m *PokemonTestMock) GetPokemon(pokemon entities.Pokemon) (interface{}, error) {
	args := m.Called(pokemon)
	return args.Get(0).(entities.Pokemon), args.Error(1)
}
