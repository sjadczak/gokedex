package pokeapi

import (
	"encoding/json"
	"fmt"

	"github.com/sjadczak/gokedex/internal/pokeapi/models"
)

func (c *Client) Pokemon(name string) (*models.Pokemon, error) {
	endpoint := fmt.Sprintf("pokemon/%s", name)

	data, err := c.do(endpoint)
	if err != nil {
		return nil, err
	}

	var pokemon models.Pokemon
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return nil, err
	}

	return &pokemon, nil
}
