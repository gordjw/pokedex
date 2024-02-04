package main

import (
	"bufio"
	"fmt"
	"os"
	"pokeapi"
	"strings"
)

type Config struct {
	prevLocationAreaURL *string
	nextLocationAreaURL *string
	pokeapiClient       pokeapi.Client
}

type command struct {
	name        string
	description string
	callback    func(*Config, []string) error
}

func main() {
	fmt.Println("Booting poke´dex")

	config := Config{
		pokeapiClient: pokeapi.NewClient(),
	}
	commands := setupCommands()
	scanner := bufio.NewScanner(os.Stdin)

	promptStr := "Pokedex > "
	fmt.Print(promptStr)

	for scanner.Scan() {
		var args []string
		command, argString, ok := strings.Cut(scanner.Text(), " ")
		if ok {
			args = strings.Split(argString, " ")
		}

		_, ok = commands[command]
		if ok {
			err := commands[command].callback(&config, args)
			if err != nil {
				fmt.Printf("Error running %s: %v\n", command, err)
			}
		} else {
			fmt.Printf("Invalid command: '%s'\n", command)
		}

		if command == "exit" {
			break
		}
		fmt.Print(promptStr)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func setupCommands() map[string]command {
	return map[string]command{
		"map": {
			name:        "map",
			description: "Display the next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Diplay the previous 20 locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Explore the current area for Pokemon",
			callback:    commandExplore,
		},
		"help": {
			name:        "help",
			description: "Displays help information",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exits the Pokedex",
			callback:    commandExit,
		},
	}
}
