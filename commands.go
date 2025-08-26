package main

import (
	"fmt"
	"os"
)

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	// for _, cmd := range cmds {
	// 	fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	// }
	fmt.Println("help: Display commands")
	fmt.Println("exit: Close the Pokedex")
	fmt.Println("map: Display the next 20 locations")
	fmt.Println("mapb: Display the previous 20 locations")
	fmt.Println()
	return nil
}
