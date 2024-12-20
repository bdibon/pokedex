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
	mapForward, mapBackward := mapCommandsFactory(20)

	pokemonStore := make(PokemonStore)
	catch := catchCommandFactory(pokemonStore)
	inspect := inspectCommandFactory(pokemonStore)
	pokedex := pokedexCommandFactory(pokemonStore)

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
		"catch": {
			"catch",
			"Catch a Pokémon to add it to your Pokédex",
			catch,
		},
		"inspect": {
			"inspect",
			"inspect a Pokémon from your Pokédex",
			inspect,
		},
		"pokedex": {
			"pokedex",
			"list all Pokémons from your Pokédex",
			pokedex,
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
