package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *config, params ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, params ...string) error {
	cmds := makeCommands()
	fmt.Println()
	fmt.Println("welcome to the pokedex!")
	fmt.Println("usage:")
	fmt.Println()
	for _, cmd := range cmds {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}
