package usecases

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"log"
	"testProject/core/entities"
	"testProject/core/usecases"
	"testProject/tests/mocks"
	"testing"
)

func TestGetPokemon(t *testing.T) {
	ctx := context.Background()

	t.Run("when_pokemon_is_found", func(t *testing.T) {
		PokemonRepoSpy := new(mocks.PokemonTestMock)

		pkm := entities.Pokemon{
			Name: "ditto",
		}

		PokemonRepoSpy.On("GetPokemon", ctx, pkm).Return(errors.New("on any error")).Run(func(args mock.Arguments) {
			arg := args.Get(0).(entities.Pokemon)
			log.Println("GetPokemon called with args: ", arg)
		})

		uc := usecases.GetPokemonUseCase(PokemonRepoSpy)
		_, err := uc.Execute(ctx, pkm)

		assert.Nil(t, err)
	})

	t.Run("when_pokemon_is_not_found", func(t *testing.T) {
		PokemonRepoSpy := new(mocks.PokemonTestMock)
		pkm := entities.Pokemon{
			Name: "dotto",
		}

		PokemonRepoSpy.On("GetPokemon", mock.Anything).Return(errors.New("on any error"))

		uc := usecases.GetPokemonUseCase(PokemonRepoSpy)
		_, err := uc.Execute(ctx, pkm)

		assert.NotNil(t, err)
	})
}
