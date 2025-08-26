package pokeapi

import (
	"encoding/json"
	"fmt"

	"github.com/sjadczak/gokedex/internal/pokeapi/models"
)

func (c *Client) LocationArea(id int) (*models.LocationArea, error) {
	endpoint := fmt.Sprintf("location-area/%d", id)

	data, err := c.do(endpoint)
	if err != nil {
		return nil, err
	}

	var location models.LocationArea
	err = json.Unmarshal(data, &location)
	if err != nil {
		return nil, err
	}

	return &location, nil
}
