package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config) error {
	pokemonsResp, err := cfg.pokeapiClient.GetPokemon(cfg.pokemonName)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonsResp.Name)
	baseXp := pokemonsResp.BaseExperience
	random := rand.Intn(baseXp)
	fmt.Println(baseXp, random)
	isCought := false
	if random < 70 {
		isCought = true
	}
	if isCought {
		fmt.Printf("%s was cought!\n", pokemonsResp.Name)
		cfg.pokedex.pokemons[pokemonsResp.Name] = pokemonsResp
		fmt.Println("You may now inspect it with the inspect command.")
	} else {
		fmt.Printf("%s escaped!\n", pokemonsResp.Name)
	}

	return nil
}
