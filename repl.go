package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type config struct{
	commands map[string]cliCommand
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config) error
}

func cleanInput(text string) []string {
	return strings.Fields(text)
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userInput := scanner.Text()
		userInputClean := cleanInput(userInput)
		supportedComand, exists := cfg.commands[strings.ToLower(userInputClean[0])]
		if exists {
			err := supportedComand.callback(cfg)
			if err != nil {
				fmt.Println("Error:", err)
			}
			continue
		}else{
			fmt.Println("Unknown command.")
			continue
		}
	}
}
