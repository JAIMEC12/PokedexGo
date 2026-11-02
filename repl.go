package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func cleanInput(text string) []string {
	return strings.Fields(text)
}

func getCommands() map[string]cliCommand{
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Show the help menu",
			callback:    commandHelp,
		},
	}
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userInput := scanner.Text()
		userInputClean := cleanInput(userInput)
		supportedComand, exists := getCommands() [strings.ToLower(userInputClean[0])]
		if exists {
			err := supportedComand.callback()
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
