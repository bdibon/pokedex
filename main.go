package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/bdibon/pokedex/commands"
)

func main() {
	commands := commands.InitCommands()

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("pokedex > ")
	for scanner.Scan() {
		text := scanner.Text()
		command, ok := commands[text]
		if ok {
			command.Callback()
		} else {
			fmt.Printf("unknown command: %s\n", text)
		}
		fmt.Print("pokedex > ")
	}
}
