package main

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/jacobhuneke/pokedex/internal/pokeapi"
)

func commandCatch(c *config, arg string) error {
	fullURL := "https://pokeapi.co/api/v2/pokemon/" + arg + "/"

	pokeBody, err := c.client.OpenFile(fullURL)
	if err != nil {
		return err
	}

	pokeData, err := pokeapi.GetPokeData(pokeBody)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokeData.Name)
	xp := pokeData.BaseExperience
	xpFloor := math.Floor(float64(xp / 100))
	num := rand.Intn(xp)

	var scalar float64
	switch xpFloor {
	case 0:
		scalar = 0.8
	case 1:
		scalar = 0.6
	case 2:
		scalar = 0.35
	case 3:
		scalar = 0.15
	default:
		scalar = 0.1
	}

	if float64(num) > (scalar * float64(xp)) {
		fmt.Printf("%s escaped!\n", pokeData.Name)
	} else {
		fmt.Printf("%s was caught!\n", pokeData.Name)
		c.pokedex[pokeData.Name] = pokeData
		fmt.Println("You may now inspect it with the inspect command")
	}

	return nil
}
