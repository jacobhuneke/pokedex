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
			var arg string
			okarg := true
			if command.name == "explore" {
				if len(strs) <= 1 {
					okarg = false
				}
			}

			if len(strs) > 1 {
				arg = strs[1]
			} else {
				arg = ""
			}

			if ok && okarg {
				command.callback(arg)
			} else if okarg {
				fmt.Println("Unknown command")
			} else {
				fmt.Println("you must provide a location")
			}
		}
	}
}
