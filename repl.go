package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/jacobhuneke/pokedex/internal/pokeapi"
	"github.com/jacobhuneke/pokedex/internal/pokecache"
)

type config struct {
	nextURL     string
	previousURL string
	client      *pokeapi.Client
}

type cliCommand struct {
	name        string
	description string
	location    string
	callback    func(arg string) error
	cfg         *config
}

func initConfig() *config {
	apiClient := pokeapi.Client{
		Cache: *pokecache.NewCache(time.Minute * 5),
	}
	c := config{
		nextURL:     "https://pokeapi.co/api/v2/location-area//",
		previousURL: "",
		client:      &apiClient,
	}
	return &c
}

func newRegistry() map[string]cliCommand {
	registry := make(map[string]cliCommand)
	c := initConfig()

	registry["exit"] = cliCommand{
		name:        "exit",
		description: "Exit the Pokedex",
		location:    "",
		callback: func(arg string) error {
			return commandExit()
		},
		cfg: c,
	}

	registry["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		location:    "",
		callback: func(arg string) error {
			return commandHelp(registry)
		},
		cfg: c,
	}

	registry["map"] = cliCommand{
		name:        "map",
		description: "Displays the names of the next 20 location areas in the Pokemon world",
		location:    "",
		callback: func(arg string) error {
			return commandMap(c)
		},
		cfg: c,
	}

	registry["mapb"] = cliCommand{
		name:        "mapb",
		description: "Displays the names of the last 20 location areas",
		location:    "",
		callback: func(arg string) error {
			return commandMapb(c)
		},
		cfg: c,
	}

	registry["explore"] = cliCommand{
		name:        "explore",
		description: "lists all pokemon in location area",
		location:    "",
		callback: func(arg string) error {
			return commandExplore(c, arg)
		},
		cfg: c,
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
