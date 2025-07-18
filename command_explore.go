package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config) error {
	if cfg.nextCommand == nil {
		return errors.New("You need to specify a location area")
	}

	locationRes, err := cfg.pokeapiClient.GetLocation(*cfg.nextCommand)
	if err != nil {
		return err
	}
	fmt.Println("Exploring... " + *cfg.nextCommand)
	for _, pok := range locationRes.Pokemon {
		fmt.Println(pok.Pokemon.Name)
	}
	return nil
}
