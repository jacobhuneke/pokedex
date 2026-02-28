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
				command.callback()
			} else {
				fmt.Println("Unknown command")
			}
		}
	}
}
