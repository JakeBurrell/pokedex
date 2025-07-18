package main

import (
	"bufio"
	"fmt"
	"github.com/JakeBurrell/pokedexcli/internal/pokeapi"
	"os"
	"strings"
)

var clientCommands map[string]cliCommand

func init() {
	clientCommands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 locations in the Pokemon world",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations in the Pokemon world",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Displays found pokemon for a specified location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempts to catch a particular pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Displays info about caught pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Displays all pokemon you have caught",
			callback:    commandPokedex,
		},
	}
}

type config struct {
	pokeapiClient        pokeapi.Client
	nextLocationsURL     *string
	previousLocationsURL *string
	nextCommand          *string
	pokedex              map[string]pokeapi.Pokemon
}

func startRepl(cfg *config) {

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			input := scanner.Text()
			words := cleanInput(input)
			command, exists := clientCommands[words[0]]
			if !exists {
				fmt.Println("Unkown Command")
				continue
			}
			if len(words) > 1 {
				cfg.nextCommand = &words[1]
			}
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading input: ", err)
		}
	}

}

func cleanInput(text string) []string {
	words := strings.ToLower(text)
	return strings.Fields(words)
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}
