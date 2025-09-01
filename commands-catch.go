package main

import (
	"errors"
	"fmt"
	"math/rand/v2"

	"github.com/sjadczak/gokedex/internal/pokeapi"
	"github.com/sjadczak/gokedex/internal/pokeapi/models"
)

func commandCatch(cfg *config, params ...string) error {
	if len(params) != 1 {
		msg := "Usage: catch [pokemon-name]\n" +
			"Example:\n\n" +
			"Pokedex > catch pikachu\n\n"

		//lint:ignore ST1005 shown to user
		return errors.New(msg)
	}

	pname := params[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", pname)
	p, err := cfg.client.Pokemon(pname)
	if errors.Is(err, pokeapi.ErrNotFound) {
		//lint:ignore ST1005 shown to user
		return errors.New("Pokemon not found... try again!")
	} else if err != nil {
		return err
	}

	if isCaught(p) {
		fmt.Printf("You caught %s!\n", p.Name)
		cfg.pokedex[p.Name] = p
	} else {
		fmt.Printf("Uh oh! %s escaped.\n", p.Name)
	}

	return nil
}

func isCaught(p *models.Pokemon) bool {
	rate := min(1/float64(p.BaseExperience)*40, 0.50)
	return rand.Float64() < rate
}
