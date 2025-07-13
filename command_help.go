package main

import (
	"fmt"
)

func commandHelp(cfg *config) error {
	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Println("Usage:\n ")
	for _, value := range clientCommands {
		fmt.Printf("%s: %s\n", value.name, value.description)
	}
	return nil
}
