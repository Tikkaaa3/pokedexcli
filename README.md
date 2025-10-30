# Pokédex CLI

A simple command line **Pokédex application** written in Go, powered by the [PokéAPI v2](https://pokeapi.co/).  
Explore locations, discover wild Pokémon, catch them, and inspect your own Pokédex all from your terminal.

---

## Features

-  **Map navigation** — explore Pokémon world locations page by page  
-  **Explore** — list Pokémon found in a selected area  
-  **Catch** — attempt to catch Pokémon based on their base stats  
-  **Pokédex** — view all Pokémon you’ve caught so far  
-  **Inspect** — view details, stats, and abilities of a caught Pokémon  
-  **Help** — see available commands and their usage  
-  **Exit** — close the application gracefully  

---

## Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/tikkaaa3/pokedexcli.git
   cd pokedexcli
   ```

2. **Build the project**
   ```bash
   go build -o pokedex
   ```

3. **Run it**
   ```bash
   ./pokedex
   ```

---

## Commands

| Command | Description |
|----------|-------------|
| `map` | Show the next set of Pokémon world locations |
| `mapb` | Show the previous set of locations |
| `explore <location>` | List Pokémon found in a given location |
| `catch <pokemon>` | Try to catch a Pokémon (chance depends on its base stats) |
| `inspect <pokemon>` | View detailed stats, types, and abilities of a caught Pokémon |
| `pokedex` | Display all Pokémon you’ve caught so far |
| `help` | Show all available commands and descriptions |
| `exit` | Exit the Pokédex CLI |

---

## Example Session

```bash
Pokedex> map
-pallet-town
-viridian-city
...

Pokedex> explore viridian-city
- pidgey
- rattata

Pokedex> catch pidgey
Throwing Pokeball at pidgey...
pidgey was cought!

Pokedex> pokedex
Your Pokedex:
- pidgey

Pokedex> inspect pidgey
Name: pidgey
Base Experience: 50
Height: 3
Weight: 18
Types:
  -normal
  -flying
Abilities:
  - keen-eye
  - tangled-feet

Pokedex> exit
Closing the Pokedex... Goodbye!
```

---

## Tech Stack

- **Language:** Go
- **API:** [PokéAPI v2](https://pokeapi.co/)
- **CLI:** Custom REPL loop for command input

---

## License

[MIT License](./LICENSE) © 2025 Emre Tolga Kaptan
Pokémon and Pokémon character names are trademarks of Nintendo, Game Freak, and The Pokémon Company.
