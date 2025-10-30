package main

import (
	"fmt"
)

func commandInspect(cfg *config) error {
	pokemon, ok := cfg.pokedex.pokemons[cfg.pokemonName]
	if ok {
		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Height: %d\n", pokemon.Height)
		fmt.Printf("Weight: %d\n", pokemon.Weight)
		fmt.Println("Stats:")
		for i := range len(pokemon.Stats) {
			fmt.Printf("\t-%s: %d\n", pokemon.Stats[i].Stat.Name, pokemon.Stats[i].BaseStat)
		}
		fmt.Println("Types:")
		for i := range len(pokemon.Types) {
			fmt.Printf("\t-%s\n", pokemon.Types[i].Type.Name)
		}
		return nil
	} else {
		fmt.Println("you have not caught that pokemon")
		return nil
	}
}
