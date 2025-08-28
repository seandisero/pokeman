package main

import (
	"fmt"
)

func commandMap(cfg *config, args ...string) error {
	locations, err := cfg.pokapiClient.GetLocationAreas(cfg.Next)
	if err != nil {
		fmt.Println(err)
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
