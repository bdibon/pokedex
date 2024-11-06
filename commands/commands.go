package commands

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	Callback    func() error
}

func InitCommands() map[string]cliCommand {
	commands := map[string]cliCommand{
		"exit": {
			"exit",
			"Exit the Pokedex REPL",
			func() error {
				os.Exit(0)
				return nil
			},
		},
		"help": {
			"help",
			"Displays a help message",
			nil,
		},
		"map": {
			"map",
			"Display the next 20 locations",
			func() error {
				// GET https://pokeapi.co/api/v2/location-area/
				// The response will contain a JSON payload with:
				//	* a 'next' field that is the next URL to query for the next page
				// 	* a 'results' field that is an array of 20 location names
				return nil
			},
		},
		"mapb": {
			"mapb",
			"Display the previous 20 locations",
			nil,
		},
	}
	initHelpCommand(commands)
	return commands
}

func initHelpCommand(commands map[string]cliCommand) {
	helpCommand := commands["help"]
	helpCommand.Callback = func() error {
		fmt.Printf("\nWelcome to the Pokedex!\nUsage:\n\n")
		for name, cmd := range commands {
			fmt.Printf("%s: %s\n", name, cmd.description)
		}
		fmt.Println()
		return nil
	}
	commands["help"] = helpCommand
}
