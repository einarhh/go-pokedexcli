package main

import (
	"fmt"
)

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex CLI!")
	fmt.Println("Available commands:")
	for _, c := range getCommands() {
		fmt.Printf("%s - %s\n", c.name, c.description)
	}
	return nil
}
