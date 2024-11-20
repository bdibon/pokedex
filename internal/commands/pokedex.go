package commands

import (
	"fmt"
	"strings"
)

func pokedexCommandFactory(pokemonStore PokemonStore) func(...string) error {
	return func(args ...string) error {
		return pokedexBaseCommand(pokemonStore)
	}
}

func pokedexBaseCommand(pokemonStore PokemonStore) error {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintln("Your Pokédex:"))
	for key := range pokemonStore {
		if _, ok := pokemonStore[key]; ok {
			sb.WriteString(fmt.Sprintf("\t- %s\n", key))
		}
	}
	fmt.Print(sb.String())
	return nil
}
