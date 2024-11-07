package commands

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/bdibon/pokedex/pokeapi"
)

const pagination = 20

var offset int = -pagination

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

func mapForward() error {
	offset += pagination
	return mapBaseCommand(offset, pagination)
}

func mapBackward() error {
	if offset < 0 {
		return errors.New("first page of results, cannot go back")
	}
	offset -= pagination
	return mapBaseCommand(offset, pagination)
}
