package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemonData(pokemonName string) (PokemonData, error) {
	url := baseURL + "/pokemon/" + pokemonName

	if val, ok := c.cache.Get(url); ok {
		locationResp := PokemonData{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return PokemonData{}, err
		}
		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonData{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonData{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonData{}, err
	}

	pokemonData := PokemonData{}
	err = json.Unmarshal(dat, &pokemonData)
	if err != nil {
		return PokemonData{}, err
	}

	return pokemonData, nil
}
