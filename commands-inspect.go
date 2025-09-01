package main

import (
	"errors"
	"fmt"

	"github.com/sjadczak/gokedex/internal/pokeapi/models"
)

func commandsInspect(cfg *config, params ...string) error {
	if len(params) != 1 {
		msg := "Usage: inspect [pokemon-name]\n" +
			"Example:\n\n" +
			"Pokedex > inspect pikachu\n\n"

		//lint:ignore ST1005 shown to user
		return errors.New(msg)
	}
	pname := params[0]
	pokemon, ok := cfg.pokedex[pname]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	printPokemon(pokemon)
	return nil
}

func printPokemon(p *models.Pokemon) {
	fmt.Printf("Name: %s\n", p.Name)
	fmt.Printf("Height: %d\n", p.Height)
	fmt.Printf("Weight: %d\n", p.Weight)
	fmt.Println("Stats:")
	for _, s := range p.Stats {
		fmt.Printf(" -%s: %v\n", s.Stat.Name, s.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range p.Types {
		fmt.Printf(" -%s\n", t.Type.Name)
	}
}
