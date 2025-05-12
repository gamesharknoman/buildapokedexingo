package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name or id")
	}
	name := args[0]

	pokemon, err := cfg.pokeapiClient.GetPokemonData(name)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	randNum := rand.Intn(pokemon.BaseExperience)

	if randNum < pokemon.BaseExperience/2 {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemon.Name)

	cfg.pokedexMap[pokemon.Name] = pokemon

	return nil
}
