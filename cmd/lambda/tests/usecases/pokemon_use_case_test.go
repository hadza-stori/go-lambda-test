package usecases

import (
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
	//ctx := context.Background()

	t.Run("when_pokemon_is_found", func(t *testing.T) {
		/* ===== Repositories mock ====== */
		PokemonRepoSpy := new(mocks.PokemonTestMock)
		pkm := entities.Pokemon{
			Name: "ditto",
		}

		PokemonRepoSpy.On("GetPokemon", mock.Anything).Return(errors.New("on any error")).Run(func(args mock.Arguments) {
			arg := args.Get(0).(entities.Pokemon)
			log.Println("GetPokemon called with args: ", arg)
		})

		PokemonRepoSpy.MethodCalled("GetPokemon", pkm)

		_, err := usecases.GetPokemon(pkm)

		assert.Nil(t, err)
		PokemonRepoSpy.AssertNumberOfCalls(t, "GetPokemon", 1)
	})

	t.Run("when_pokemon_is_not_found", func(t *testing.T) {
		/* ===== Repositories mock ====== */
		PokemonRepoSpy := new(mocks.PokemonTestMock)
		pkm := entities.Pokemon{
			Name: "dotto",
		}

		PokemonRepoSpy.On("GetPokemon", mock.Anything).Return(errors.New("on any error")).Run(func(args mock.Arguments) {
			log.Println("GetPokemon called with args: ", args)
		})

		PokemonRepoSpy.MethodCalled("GetPokemon", pkm)

		_, err := usecases.GetPokemon(pkm)

		assert.NotNil(t, err)
		PokemonRepoSpy.AssertNumberOfCalls(t, "GetPokemon", 1)
	})
}
