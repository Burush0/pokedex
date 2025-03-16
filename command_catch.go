package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]

	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	if isSuccessful(pokemon.BaseExperience) {
		fmt.Printf("%s was caught!\n", pokemon.Name)
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}

	cfg.caughtPokemon[pokemon.Name] = pokemon

	return nil
}

// Function to calculate the chance of success
func calculateChanceOfSuccess(number int) float64 {
	// Normalize the number to a range of 0 to 1
	// Lower number (50) should give higher chance (1), higher number (200) should give lower chance (0)
	normalized := 1.0 - float64(number-50)/150.0

	// Ensure the normalized value is within [0, 1]
	if normalized < 0 {
		normalized = 0
	} else if normalized > 1 {
		normalized = 1
	}

	return normalized
}

// Function to determine if the action is successful
func isSuccessful(number int) bool {

	// Calculate the chance of success
	chance := calculateChanceOfSuccess(number)

	// Generate a random number between 0 and 1
	randomValue := rand.Float64()

	// Determine success
	return randomValue <= chance
}
