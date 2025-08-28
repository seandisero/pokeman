package main

import (
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("error: no arguments passed into explore commane")
	}
	locationData, err := cfg.pokapiClient.GetLocationData(args[0])
	if err != nil {
		return fmt.Errorf("error getting location data: %w", err)
	}

	fmt.Printf("Exploring %s...\n", args[0])
	fmt.Println("Found Pokemon: ")
	for _, v := range locationData.PokemonEncounters {
		if v.Pokemon.Name != nil {
			fmt.Printf(" - %s\n", *v.Pokemon.Name)
		}
	}

	return nil
}
