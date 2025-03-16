package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.caughtPokemon) == 0 {
		return errors.New("you have no caught pokemon")
	}

	fmt.Println("Your Pokedex:")
	for key := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", key)
	}

	return nil
}
