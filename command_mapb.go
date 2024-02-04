package main

import "fmt"

func commandMapb(config *Config, args []string) error {
	locationAreaList, err := config.pokeapiClient.LocationAreas(config.prevLocationAreaURL)
	if err != nil {
		return err
	}

	fmt.Println("Next 20 locations:")
	for i := 0; i < len(locationAreaList.Areas); i++ {
		fmt.Println(locationAreaList.Areas[i].Name)
	}

	config.nextLocationAreaURL = locationAreaList.Next
	config.prevLocationAreaURL = locationAreaList.Previous

	return nil
}
