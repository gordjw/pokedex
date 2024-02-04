package pokeapi

import (
	"encoding/json"
	"fmt"
)

func (c *Client) Locations(url string) (LocationList, error) {
	var fullUrl string
	if url == "" {
		fullUrl = baseUrl + "/location"
	} else {
		fullUrl = url
	}

	bytes, err := c.getPokemonApiData(fullUrl)
	if err != nil {
		return LocationList{}, err
	}

	var locationList LocationList
	err = json.Unmarshal(bytes, &locationList)
	if err != nil {
		return LocationList{}, err
	}

	return locationList, nil
}

func (c *Client) Location(id int) (Location, error) {
	bytes, err := c.getPokemonApiData(fmt.Sprintf("/location/%d", id))
	if err != nil {
		return Location{}, err
	}

	var location Location
	err = json.Unmarshal(bytes, &location)
	if err != nil {
		return Location{}, err
	}

	return location, nil
}
