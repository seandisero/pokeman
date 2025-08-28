package main

import (
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	// display name height weight stats: hp attack defence special-attack spacial-defence speed as well as the types

	if len(args) < 1 {
		return fmt.Errorf("no args passed into inspect")
	}

	pokemon, exists := cfg.pokedex[args[0]]
	if !exists {
		return fmt.Errorf("could not find pokemon %s", args[0])
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, s := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", *s.Stat.Name, s.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("  - %s\n", *t.Type.Name)
	}

	return nil
}
