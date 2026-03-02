package main

import (
	"fmt"

	"github.com/jacobhuneke/pokedex/internal/pokeapi"
)

func commandExplore(c *config, location string) error {
	fullURL := "https://pokeapi.co/api/v2/location-area/" + location + "/"
	locBody, err := c.client.OpenFile(fullURL)
	if err != nil {
		return err
	}

	locdata, err := pokeapi.GetDataLoc(locBody)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", location)
	fmt.Println("Found Pokemon:")
	for pokemon := range locdata.PokemonEncounters {
		fmt.Printf(" - %v\n", locdata.PokemonEncounters[pokemon].Pokemon.Name)
	}
	return nil
}
