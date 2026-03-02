package main

import (
	"errors"
	"fmt"
)

func commandInspect(c *config, arg string) error {
	val, ok := c.pokedex[arg]

	if !ok {
		fmt.Println("you have not caught that pokemon")
		return errors.New("you have not caught that pokemon")
	} else {
		fmt.Printf("Name: %s\n", val.Name)
		fmt.Printf("Height: %v\n", val.Height)
		fmt.Printf("Weight: %v\n", val.Weight)
		fmt.Printf("Stats: \n")
		for i := range val.Stats {
			fmt.Printf(" -%v: %v\n", val.Stats[i].Stat.Name, val.Stats[i].BaseStat)
		}
		fmt.Printf("Types: \n")
		for i := range val.Types {
			fmt.Printf(" - %v\n", val.Types[i].Type.Name)
		}
	}

	return nil
}
