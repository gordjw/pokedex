package pokeapi

import (
	"encoding/json"
	"fmt"
)

func (c *Client) LocationAreas(url *string) (LocationAreaList, error) {
	endpoint := "/location-area"
	fullUrl := baseUrl + endpoint
	if url != nil {
		fullUrl = *url
	}

	bytes, err := c.getPokemonApiData(fullUrl)
	if err != nil {
		return LocationAreaList{}, err
	}

	var areaList LocationAreaList
	err = json.Unmarshal(bytes, &areaList)
	if err != nil {
		return LocationAreaList{}, err
	}

	return areaList, nil
}

func (c *Client) LocationArea(id int) (Area, error) {
	endpoint := fmt.Sprintf("/location-area/%d", id)
	fullUrl := baseUrl + endpoint

	bytes, err := c.getPokemonApiData(fullUrl)
	if err != nil {
		return Area{}, err
	}

	var area Area
	err = json.Unmarshal(bytes, &area)
	if err != nil {
		return Area{}, err
	}

	return area, nil
}
