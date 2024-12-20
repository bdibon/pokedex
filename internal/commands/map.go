package commands

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/bdibon/pokedex/internal/pokeapi"
)

func mapCommandsFactory(pagination int) (func(...string) error, func(...string) error) {
	offset := -pagination

	mapForward := func(_ ...string) error {
		offset += pagination
		return mapBaseCommand(offset, pagination)
	}
	mapBackward := func(_ ...string) error {
		if offset < 20 {
			return errors.New("first page of results, cannot go back")
		}
		offset -= pagination
		return mapBaseCommand(offset, pagination)
	}
	return mapForward, mapBackward
}

func mapBaseCommand(of, pag int) error {
	res, err := pokeapi.GetLocationAreas(of, pag)
	if err != nil {
		return fmt.Errorf("error getting location areas: %w", err)
	}

	writer := bufio.NewWriter(os.Stdout)
	for _, r := range res {
		writer.WriteString(fmt.Sprintln(r.Name))
	}
	writer.Flush()

	return nil
}
