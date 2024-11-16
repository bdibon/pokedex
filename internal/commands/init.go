package commands

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	Callback    func(arguments ...string) error
}

func InitCommands() map[string]cliCommand {
	commands := map[string]cliCommand{
		"exit": {
			"exit",
			"Exit the Pokedex REPL",
			func(_ ...string) error {
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
			mapForward,
		},
		"mapb": {
			"mapb",
			"Display the previous 20 locations",
			mapBackward,
		},
		"explore": {
			"explore",
			"See a list of all the Pokémon in a given area",
			explore,
		},
	}
	initHelpCommand(commands)
	return commands
}

func initHelpCommand(commands map[string]cliCommand) {
	helpCommand := commands["help"]
	helpCommand.Callback = func(_ ...string) error {
		fmt.Printf("\nWelcome to the Pokedex!\nUsage:\n\n")
		for name, cmd := range commands {
			fmt.Printf("%s: %s\n", name, cmd.description)
		}
		fmt.Println()
		return nil
	}
	commands["help"] = helpCommand
}
