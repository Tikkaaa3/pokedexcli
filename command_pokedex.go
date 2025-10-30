package main

import (
	"fmt"
)

func commandPokedex(cfg *config) error {
	pokedex := cfg.pokedex.pokemons
	fmt.Println("Your Pokedex:")
	for _, v := range pokedex {
		fmt.Printf("\t-%s\n", v.Name)
	}
	return nil
}
