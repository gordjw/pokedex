package main

import "fmt"

func commandExit(config *Config, args []string) error {
	fmt.Println("Exiting")
	return nil
}
