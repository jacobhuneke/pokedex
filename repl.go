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
	pokedex     map[string]pokeapi.PokemonData
}

type cliCommand struct {
	name        string
	description string
	arg         string
	callback    func(arg string) error
	cfg         *config
}

func readMD() []byte {
	data, err := os.ReadFile("user.md")
	if err != nil {
		return nil
	}
	return data
}

func initConfig() (*config, error) {
	apiClient := pokeapi.Client{
		Cache: *pokecache.NewCache(time.Minute * 5),
	}
	pokedex := make(map[string]pokeapi.PokemonData)
	bytes := readMD()
	if bytes != nil {
		str := string(bytes)
		pokemon := strings.Split(str, "\n")
		for _, p := range pokemon {
			if p != "" && p != "\n" {
				fullURL := "https://pokeapi.co/api/v2/pokemon/" + p + "/"
				pokeBody, err := apiClient.OpenFile(fullURL)
				if err != nil {
					return nil, err
				}

				pokeData, err := pokeapi.GetPokeData(pokeBody)
				if err != nil {
					return nil, err
				}
				pokedex[p] = pokeData
			}
		}
	}
	c := config{
		nextURL:     "https://pokeapi.co/api/v2/location-area//",
		previousURL: "",
		client:      &apiClient,
		pokedex:     pokedex,
	}
	return &c, nil
}

func newRegistry() map[string]cliCommand {
	registry := make(map[string]cliCommand)
	c, ok := initConfig()
	if ok != nil {
		return nil
	}

	registry["exit"] = cliCommand{
		name:        "exit",
		description: "Exit the Pokedex",
		arg:         "",
		callback: func(arg string) error {
			return commandExit(c)
		},
		cfg: c,
	}

	registry["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		arg:         "",
		callback: func(arg string) error {
			return commandHelp(registry)
		},
		cfg: c,
	}

	registry["map"] = cliCommand{
		name:        "map",
		description: "Displays the names of the next 20 location areas in the Pokemon world",
		arg:         "",
		callback: func(arg string) error {
			return commandMap(c)
		},
		cfg: c,
	}

	registry["mapb"] = cliCommand{
		name:        "mapb",
		description: "Displays the names of the last 20 location areas",
		arg:         "",
		callback: func(arg string) error {
			return commandMapb(c)
		},
		cfg: c,
	}

	registry["explore"] = cliCommand{
		name:        "explore",
		description: "lists all pokemon in location area",
		arg:         "",
		callback: func(arg string) error {
			return commandExplore(c, arg)
		},
		cfg: c,
	}

	registry["catch"] = cliCommand{
		name:        "catch",
		description: "Tries to catch the named Pokemon",
		arg:         "",
		callback: func(arg string) error {
			return commandCatch(c, arg)
		},
		cfg: c,
	}

	registry["inspect"] = cliCommand{
		name:        "inspect",
		description: "prints info about a caught pokemon",
		arg:         "",
		callback: func(arg string) error {
			return commandInspect(c, arg)
		},
		cfg: c,
	}

	registry["pokedex"] = cliCommand{
		name:        "pokedex",
		description: "prints a list of all Pokemon you have caught",
		arg:         "",
		callback: func(arg string) error {
			return commandPokedex(c)
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

func commandExit(c *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")

	pokemon := []string{}
	for poke := range c.pokedex {
		pokemon = append(pokemon, poke)
	}
	data := []byte(strings.Join(pokemon, "\n"))
	err := os.WriteFile("user.md", data, 0644)
	if err != nil {
		return err
	}
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

func commandPokedex(c *config) error {
	fmt.Println("Your Pokedex:")
	for key := range c.pokedex {
		fmt.Printf(" - %s\n", key)
	}
	return nil
}
