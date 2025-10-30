package main

import (
	"fmt"
)

func commandExplore(cfg *config) error {
	pokemonsResp, err := cfg.pokeapiClient.ListPokemons(cfg.areaName)
	if err != nil {
		return err
	}

	for _, poke := range pokemonsResp {
		fmt.Println(poke)
	}
	return nil
}
