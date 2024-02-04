package main

import (
	"errors"
	"fmt"
)

func commandHelp(config *Config, args []string) error {
	fmt.Println("Help content goes here")
	return errors.New("Couldn't find anything to help with")
}
