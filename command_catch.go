package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args []string) error {
	if len(args) != 1 {
		return errors.New("Pokemon name not provided")
	}
	pokemonName := args[0]

	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	const threshold = 40
	randomInt := rand.Intn(pokemon.BaseExperience)

	fmt.Println(randomInt)
	if randomInt > threshold {
		return fmt.Errorf("%s escaped!", pokemonName)
	}

	fmt.Printf("%s caught!\n", pokemonName)
	cfg.pokemon[pokemonName] = pokemon

	return nil
}
