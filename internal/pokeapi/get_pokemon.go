package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	url := "https://pokeapi.co/api/v2/pokemon/" + pokemonName

	if val, ok := c.cache.Get(url); ok {
		var pokeResp Pokemon
		err := json.Unmarshal(val, &pokeResp)
		if err != nil {
			return Pokemon{}, err
		}

		return pokeResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokeResp := Pokemon{}
	err = json.Unmarshal(dat, &pokeResp)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, dat)
	return pokeResp, nil
}
