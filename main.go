package main

import (
	"github.com/sjadczak/gokedex/internal/pokeapi"
)

func main() {
	client := pokeapi.NewClient()
	startRepl(client)
}
