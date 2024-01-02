package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args []string) error {
	fmt.Println("Pokemons:")
	for _, pokemon := range cfg.pokemon {
		fmt.Printf("- %s\n", pokemon.Name)
	}

	return nil
}
