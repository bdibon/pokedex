package commands

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/bdibon/pokedex/internal/pokeapi"
)

const maxBaseExperience = 608
const minBaseExperience = 20

type PokemonStore map[string]pokeapi.Pokemon

func catchCommandFactory(pokemonStore PokemonStore) func(args ...string) error {
	catchCmd := func(args ...string) error {
		return catchBaseCommand(pokemonStore, args...)
	}
	return catchCmd
}

func catchBaseCommand(pokemonStore PokemonStore, args ...string) error {
	if len(args) < 1 {
		return errors.New("catch command: can't catch the air")
	}
	name := args[0]
	pokemon, err := pokeapi.GetPokemon(name)
	if err != nil {
		var pkmnApiErr pokeapi.PokeApiError
		if errors.As(err, &pkmnApiErr); pkmnApiErr.Code == pokeapi.ResourceNotFound {
			return fmt.Errorf("pokemon `%s` doesn't exist", name)
		}
		return err
	}
	fmt.Printf("Throwing a pokeball at %s...\n", name)
	if outcome := throwPokeball(pokemon); outcome {
		fmt.Printf("%s was caught!\n", name)
		pokemonStore[name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", name)
	}
	return nil
}

func throwPokeball(pokemon pokeapi.Pokemon) bool {
	random := 20 + rand.Intn(maxBaseExperience+1)
	return random > pokemon.BaseExperience
}
