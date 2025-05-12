package main

import (
	"errors"
	"fmt"

	"github.com/gamesharknoman/pokedexcli/internal/pokeapi"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name or id")
	}
	name := args[0]
	_, exists := cfg.pokedexMap[name]
	if !exists {
		return errors.New("you have not caught that pokemon")
	}
	pokemon, err := cfg.pokeapiClient.GetPokemonData(name)
	if err != nil {
		return err
	}
	printPokemonData(pokemon)
	return nil
}

func printPokemonData(pokemon pokeapi.PokemonData) {
	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)
	fmt.Println("Stats:")
	for _, value := range pokemon.Stats {
		fmt.Printf(" -%s: %v\n", value.Stat.Name, value.BaseStat)
	}
	fmt.Println("Types:")
	for _, value := range pokemon.Types {
		fmt.Printf(" - %s\n", value.Type.Name)
	}
}
