package commands

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/bdibon/pokedex/internal/pokeapi"
)

func explore(args ...string) error {
	if len(args) == 0 {
		return errors.New("explore command: missing argument")
	}
	pokemonEncounters, err := pokeapi.GetPokemonEncounters(args[0])
	if err != nil {
		return err
	}

	writer := bufio.NewWriter(os.Stdout)
	for _, pe := range pokemonEncounters {
		io.WriteString(writer, fmt.Sprintln(pe.Pokemon.Name))
	}
	writer.Flush()

	return nil
}
