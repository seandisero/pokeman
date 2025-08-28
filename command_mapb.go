package main

import (
	"errors"
	"fmt"
)

func commandMapb(cfg *config, args ...string) error {
	locations, err := cfg.pokapiClient.GetLocationAreas(cfg.Previous)
	if err != nil {
		return errors.New("You are on the first page")
	}

	if locations.Next != nil {
		cfg.Next = locations.Next
	}
	if locations.Previous != nil {
		cfg.Previous = locations.Previous
	}

	for _, l := range locations.Results {
		fmt.Println(l.Name)
	}
	return nil
}
