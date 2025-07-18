package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config) error {
	if cfg.nextCommand == nil {
		return errors.New("You need to specify a pokemon to inspect")
	}
	pokemon, exists := cfg.pokedex[*cfg.nextCommand]
	if !exists {
		return errors.New("You have not caught this pokemon yet")
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)

	fmt.Printf("Stats:\n")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" -%s: %d\n", stat.Stat.Name, stat.Lvl)
	}

	fmt.Println("Types:")
	for _, typ := range pokemon.Types {
		fmt.Printf(" - %s\n", typ.Typ.Name)
	}

	return nil
}
