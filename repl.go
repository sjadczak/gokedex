package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/sjadczak/gokedex/internal/pokeapi"
)

func startRepl(client *pokeapi.Client) {
	cmds := makeCommands(client)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		command, ok := cmds[commandName]
		if ok {
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	words := strings.Fields(lower)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func makeCommands(client *pokeapi.Client) map[string]cliCommand {
	cm, cmb := makeMapCommands(client)

	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Show next 20 locations",
			callback:    cm,
		},
		"mapb": {
			name:        "mapb",
			description: "Show previous 20 locations",
			callback:    cmb,
		},
	}
}
