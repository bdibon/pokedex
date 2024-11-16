package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/bdibon/pokedex/internal/commands"
)

func main() {
	cmds := commands.InitCommands()

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("pokedex > ")
	for scanner.Scan() {
		text := scanner.Text()
		args := strings.Fields(text)
		if len(args) == 0 {
			continue
		}
		cmdName := args[0]
		cmdArgs := args[1:]
		cmd, ok := cmds[cmdName]
		if ok {
			err := cmd.Callback(cmdArgs...)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %s\n", err)
			}
		} else {
			fmt.Printf("unknown command: %s\n", text)
		}
		fmt.Print("pokedex > ")
	}
}
