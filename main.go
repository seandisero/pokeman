package main

import (
	"time"

	"pokeman/internal/pokapi"
)

func main() {
	pokClient := pokapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config{
		pokapiClient: pokClient,
		pokedex:      make(map[string]pokapi.PokemonData),
	}

	startRepl(cfg)
}
