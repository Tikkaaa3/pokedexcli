package main

import (
	"github.com/Tikkaaa3/pokedexcli/internal/pokeapi"
	"time"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	pokedex := &Pokedex{
		pokemons: make(map[string]pokeapi.Pokemon),
	}
	cfg := &config{
		pokedex:       pokedex,
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}
