package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *config, parameters []string) error {
	resp, err := cfg.pokeapiClient.ListLocationArea(cfg.nextLocationAreaUrl)
	if err != nil {
		return err
	}
	for _, location := range resp.Results {
		fmt.Println(location.Name)
	}
	cfg.nextLocationAreaUrl = resp.Next
	cfg.prevLocationAreaUrl = resp.Previous
	return nil
}

func commandMapb(cfg *config, parameters []string) error {
	if cfg.prevLocationAreaUrl == nil {
		return errors.New("You are already at the beginning of the list")
	}
	resp, err := cfg.pokeapiClient.ListLocationArea(cfg.prevLocationAreaUrl)
	if err != nil {
		return err
	}
	for _, location := range resp.Results {
		fmt.Println(location.Name)
	}
	cfg.nextLocationAreaUrl = resp.Next
	cfg.prevLocationAreaUrl = resp.Previous
	return nil
}
