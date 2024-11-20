package commands

import (
	"errors"
	"fmt"
	"strings"

	"github.com/bdibon/pokedex/internal/pokeapi"
)

type Pokemon pokeapi.Pokemon

func (p Pokemon) String() string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("Name: %s\n", p.Name))
	sb.WriteString(fmt.Sprintf("Height: %d\n", p.Height))

	sb.WriteString(fmt.Sprintln("Stats:"))
	for _, statItem := range p.Stats {
		sb.WriteString(fmt.Sprintf("\t-%s: %d\n", statItem.Stat.Name, statItem.BaseStat))
	}

	sb.WriteString(fmt.Sprintln("Types:"))
	for _, typeItem := range p.Types {
		sb.WriteString(fmt.Sprintf("\t- %s\n", typeItem.Type.Name))
	}
	return sb.String()
}

func inspectCommandFactory(pokedex PokemonStore) func(...string) error {
	return func(args ...string) error {
		return inspectBaseCommand(pokedex, args...)
	}
}

func inspectBaseCommand(pokedex PokemonStore, args ...string) error {
	if len(args) == 0 {
		return errors.New("missing pokemon name")
	}

	pokemon, ok := pokedex[args[0]]
	if ok == false {
		return fmt.Errorf("you have not caught %s", args[0])
	}
	fmt.Print(Pokemon(pokemon))
	return nil
}
