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
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Show help",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the program",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Show the nex 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Show the previous 20 locations",
			callback:    commandMapb,
		},
	}
}

func cleanInput(input string) []string {
	input = strings.ToLower(input)
	return strings.Fields(input)
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}

		commandName := input[0]

		availableCommands := getCommands()

		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Printf("Command '%s' not found\n", input)
			continue
		}

		err := command.callback(cfg)
		if err != nil {
			fmt.Println(err)
		}
	}
}
