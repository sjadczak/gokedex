package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/sjadczak/gokedex/internal/pokeapi"
)

func startRepl() {
	cfg := &config{
		client: pokeapi.NewClient(),
		ls:     newLState(),
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		params := words[1:]
		command, ok := makeCommands()[commandName]
		if ok {
			err := command.callback(cfg, params...)
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

type config struct {
	client *pokeapi.Client
	ls     *locationState
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, params ...string) error
}

func makeCommands() map[string]cliCommand {
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
			description: "Show next page of locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Show previous page of locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Show pokemon found in a location",
			callback:    commandExplore,
		},
	}
}
