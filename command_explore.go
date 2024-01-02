package main

import (
	"fmt"
)

func commandExplore(cfg *config, parameters []string) error {
	locationAreaName := parameters[0]
	fmt.Printf("Exploring %s...\n", locationAreaName)
	resp, err := cfg.pokeapiClient.ExploreLocationArea(locationAreaName)
	if err != nil {
		return err
	}
	if len(resp.PokemonEncounters) > 0 {
		fmt.Println("Found pokemon:")
	}
	for _, encounter := range resp.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}
	return nil
}
