package pokapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemonData(name string) (PokemonData, error) {
	pokemonURL := basePokemonURL + "/" + name

	if pokemonRawData, exists := c.cache.Get(pokemonURL); exists {
		pokemonData := PokemonData{}
		err := json.Unmarshal(pokemonRawData, &pokemonData)
		if err != nil {
			return PokemonData{}, fmt.Errorf("error unmarshaling cached data: %w ", err)
		}
		return pokemonData, nil
	}

	req, err := http.NewRequest("GET", pokemonURL, nil)
	if err != nil {
		return PokemonData{}, fmt.Errorf("error making new request: %w ", err)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonData{}, fmt.Errorf("error performing request: %w ", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return PokemonData{}, fmt.Errorf("error reading body: %w ", err)
	}

	pokemonData := PokemonData{}
	err = json.Unmarshal(body, &pokemonData)
	if err != nil {
		return PokemonData{}, fmt.Errorf("error unmarshaling cached data: %w ", err)
	}

	c.cache.Add(pokemonURL, body)
	return pokemonData, nil
}
