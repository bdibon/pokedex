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
			err := command.Callback()
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %s\n", err)
			}
		} else {
			fmt.Printf("unknown command: %s\n", text)
		}
		fmt.Print("pokedex > ")
	}
}
