package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListPokemons(areaName string) ([]string, error) {
	url := baseURL + "/location-area" + "/" + areaName

	if val, ok := c.cache.Get(url); ok {
		var pokeResp []string
		err := json.Unmarshal(val, &pokeResp)
		if err != nil {
			return nil, err
		}

		return pokeResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	areaResp := RespArea{}
	err = json.Unmarshal(dat, &areaResp)
	if err != nil {
		return nil, err
	}

	var pokemons []string
	for _, enc := range areaResp.PokemonEncounters {
		pokemons = append(pokemons, enc.Pokemon.Name)
	}

	val, err := json.Marshal(pokemons)
	c.cache.Add(url, val)
	return pokemons, nil
}
