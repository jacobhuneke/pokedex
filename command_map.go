package main

import (
	"fmt"

	"github.com/jacobhuneke/pokedex/internal/pokeapi"
)

func commandMap(c *config) error {
	body, err := pokeapi.OpenFile(c.nextURL)
	if err != nil {
		return err
	}

	data, err := pokeapi.GetDataLoc(body)
	if err != nil {
		return err
	}

	for i := range 20 {
		fmt.Println(data.Results[i].Name)
	}
	//update urls
	if prev, ok := data.Previous.(string); ok {
		c.previousURL = prev
	} else {
		c.previousURL = ""
	}
	c.nextURL = data.Next
	return nil
}

func commandMapb(c *config) error {
	if c.previousURL == "" || c.previousURL == "null" {
		fmt.Println("you're on the first page")
		return nil
	}

	body, err := pokeapi.OpenFile(c.nextURL)
	if err != nil {
		return err
	}

	data, err := pokeapi.GetDataLoc(body)
	if err != nil {
		return err
	}

	for i := range 20 {
		fmt.Println(data.Results[i].Name)
	}

	c.nextURL = data.Next
	if prev, ok := data.Previous.(string); ok {
		c.previousURL = prev
	} else {
		c.previousURL = ""
	}

	return nil
}
