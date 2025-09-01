package main

import (
	"errors"
	"fmt"

	"github.com/sjadczak/gokedex/internal/pokeapi"
)

func commandExplore(cfg *config, params ...string) error {
	if len(params) != 1 {
		msg := "Usage: explore [location-name]\n" +
			"Example:\n\n" +
			"Pokedex > explore pastoria-city-area\n\n"

		//lint:ignore ST1005 shown to user
		return errors.New(msg)
	}

	fmt.Printf("Exploring %s...\n", params[0])
	loc, err := cfg.client.LocationArea(params[0])
	if errors.Is(err, pokeapi.ErrNotFound) {
		//lint:ignore ST1005 shown to user
		return errors.New("Location not found... try again!")
	} else if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, pe := range loc.PokemonEncounters {
		fmt.Printf(" - %s\n", pe.Pokemon.Name)
	}

	return nil
}
