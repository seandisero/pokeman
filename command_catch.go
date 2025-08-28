package main

import (
	"fmt"
	"math/rand"
	"time"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("error: no args passed into catch command")
	}

	fmt.Printf("Throwing a Pokeball at %s", args[0])

	pokemon, err := cfg.pokapiClient.GetPokemonData(args[0])
	if err != nil {
		return err
	}

	chance := rand.Int() % 300
	exp := pokemon.BaseExperience

	timer := 0.0

	for timer < float64(exp)/10.0 {
		time.Sleep(300 * time.Millisecond)
		fmt.Printf(".")
		timer += 1.0
	}

	fmt.Println()

	if chance > exp {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		cfg.pokedex[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}

	return nil
}
