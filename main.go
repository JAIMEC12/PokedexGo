package main

func main() {
	cfg := &config{
		commands: map[string]cliCommand{
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
		},
	}
	startRepl(cfg)
}
