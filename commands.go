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

func commandPokedex(cfg *config, params ...string) error {
	if len(cfg.pokedex) == 0 {
		fmt.Println("you're Pokedex is empty, go catch some pokemon!")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for pokemon := range cfg.pokedex {
		fmt.Printf(" -%s\n", pokemon)
	}
	return nil
}
