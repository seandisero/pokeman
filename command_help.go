package main

import (
	"fmt"
)

func commandHelp(c *config, args ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for _, v := range getCommands() {
		fmt.Printf("%s: %s\n", v.name, v.description)
	}
	return nil
}
