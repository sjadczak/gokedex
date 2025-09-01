package pokeapi

import (
	"encoding/json"
	"fmt"

	"github.com/sjadczak/gokedex/internal/pokeapi/models"
)

func (c *Client) LocationList(endpoint string) (*models.LocationList, error) {
	u, err := preprocessEndpoint(endpoint)
	if err != nil {
		return nil, fmt.Errorf("pokeapi.Client.LocationList - could not trim endpoint: %w", err)
	}

	data, err := c.do(u.String())
	if err != nil {
		return nil, err
	}

	var locations models.LocationList
	err = json.Unmarshal(data, &locations)
	if err != nil {
		return nil, err
	}

	return &locations, nil
}

func (c *Client) LocationArea(id string) (*models.LocationArea, error) {
	endpoint := fmt.Sprintf("location-area/%s", id)

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
