package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	registry := newRegistry()

	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			str := scanner.Text()
			strs := cleanInput(str)
			command, ok := registry[strs[0]]
			if ok {
				if len(strs) > 1 {
					arg := strs[1]
					command.callback(arg)
				} else if command.name == "explore" {
					fmt.Println("you must provide a location")
				} else if command.name == "catch" {
					fmt.Println("you must provide a pokemon")
				} else if command.name == "inspect" {
					fmt.Println("you must provide a pokemon")
				} else {
					command.callback("")
				}
			} else {
				fmt.Println("Unknown command")
			}
		}
	}
}
