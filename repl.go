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
	}
}

type config struct {
	pokeapiClient        pokeapi.Client
	nextLocationsURL     *string
	previousLocationsURL *string
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
