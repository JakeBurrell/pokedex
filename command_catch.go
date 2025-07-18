package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config) error {
	if cfg.nextCommand == nil {
		return errors.New("Youu need to specify a pokemon to catch")
	}

	pokemonRes, err := cfg.pokeapiClient.GetPokemon(*cfg.nextCommand)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonRes.Name)
	if (rand.Intn(10) * pokemonRes.Base_experience) < 5 {
		fmt.Printf("%s was caught!\n", pokemonRes.Name)
		return nil
	}
	fmt.Printf("%s excaped!\n", pokemonRes.Name)
	return nil

}
