package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/bdibon/pokedex/internal/commands"
)

func main() {
	cmds := commands.InitCommands()

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("pokedex > ")
	for scanner.Scan() {
		text := scanner.Text()
		cmd, ok := cmds[text]
		if ok {
			err := cmd.Callback()
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %s\n", err)
			}
		} else {
			fmt.Printf("unknown command: %s\n", text)
		}
		fmt.Print("pokedex > ")
	}
}
