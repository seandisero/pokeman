package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	fmt.Println("Pokedex: ")
	for k := range cfg.pokedex {
		fmt.Printf(" - %s\n", k)
	}
	return nil
}
