package main

import (
	"time"

	"github.com/einarhh/go-pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationAreaUrl *string
	prevLocationAreaUrl *string
	pokemon             map[string]pokeapi.Pokemon
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
		pokemon:       make(map[string]pokeapi.Pokemon),
	}
	startRepl(&cfg)
}
