package main

import (
	"bufio"
	"fmt"
	"os"
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
			description: "Displays the names of the top 20 locations in the Pokemon world",
			callback:    nil,
		},
	}

}

func main() {

	fmt.Println("Welcome to the Pokedex!")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			input := scanner.Text()
			words := cleanInput(input)
			command, ok := clientCommands[words[0]]
			if !ok {
				fmt.Println("Unkown Command")
				continue
			}
			command.callback()
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading input: ", err)
		}
	}
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Usage:\n ")
	for _, value := range clientCommands {
		fmt.Printf("%s: %s\n", value.name, value.description)
	}
	return nil
}

func commandMap() error {
	return nil
}

type config struct {
	Next     string
	Previous string
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}
