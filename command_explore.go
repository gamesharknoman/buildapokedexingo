package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}
	name := args[0]
	location, err := cfg.pokeapiClient.GetLocationData(name)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %v...\n", args)
	fmt.Println("Found Pokemon:")
	for _, encounter := range location.PokemonEncounters {
		fmt.Println(" - ", encounter.Pokemon.Name)
	}
	return nil
}
