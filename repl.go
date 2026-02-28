package main

import (
	"fmt"
	"os"
	"strings"
)

type config struct {
	nextURL     string
	previousURL string
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
	location    *config
}

func initConfig() config {
	c := config{
		nextURL:     "https://pokeapi.co/api/v2/location-area//",
		previousURL: "",
	}
	return c
}

func newRegistry() map[string]cliCommand {
	registry := make(map[string]cliCommand)
	c := initConfig()

	registry["exit"] = cliCommand{
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
		location:    &c,
	}

	registry["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback: func() error {
			return commandHelp(registry)
		},
		location: &c,
	}

	registry["map"] = cliCommand{
		name:        "map",
		description: "Displays the names of the next 20 location areas in the Pokemon world",
		callback: func() error {
			return commandMap(&c)
		},
		location: &c,
	}

	registry["mapb"] = cliCommand{
		name:        "mapb",
		description: "Displays the names of the last 20 location areas",
		callback: func() error {
			return commandMapb(&c)
		},
		location: &c,
	}

	return registry
}

func cleanInput(text string) []string {
	var strs []string
	lower := strings.ToLower(text)
	strs = strings.Fields(lower)
	return strs
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(registry map[string]cliCommand) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")

	for cmd := range registry {
		str := cmd + ": " + registry[cmd].description
		fmt.Println(str)
	}
	return nil
}
