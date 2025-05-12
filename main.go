package main

import (
	"time"

	"github.com/gamesharknoman/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	cfg := &config{
		pokedexMap:    map[string]pokeapi.PokemonData{},
		pokeapiClient: pokeClient,
	}
	startRepl(cfg)
}
