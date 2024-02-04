package main

import (
	"errors"
	"fmt"
	"strconv"
)

func commandExplore(config *Config, args []string) error {
	if len(args) == 0 {
		return errors.New("Too few arguments")
	}
	id, err := strconv.Atoi(args[0])
	if err != nil {
		return errors.New(fmt.Sprintf("Invalid argument: %v", args[0]))
	}

	locationArea, err := config.pokeapiClient.LocationArea(id)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s\n\n", locationArea.Name)
	for i := 0; i < len(locationArea.Pokemon_encounters); i++ {
		fmt.Println(locationArea.Pokemon_encounters[i].Pokemon.Name)
	}
	return nil
}
