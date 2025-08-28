package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"pokeman/internal/pokapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(c *config, args ...string) error
}

type config struct {
	pokapiClient pokapi.Client
	pokedex      map[string]pokapi.PokemonData
	Next         *string
	Previous     *string
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("Pokedex > ")
		if !scanner.Scan() {
			break
		}
		wordList := cleanInput(scanner.Text())
		if len(wordList) < 1 {
			continue
		}

		commandName := wordList[0]
		args := []string{}
		if len(wordList) > 1 {
			args = wordList[1:]
		}
		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func getCommands() map[string]cliCommand {
	commands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"pokedex": {
			name:        "pokedex",
			description: "display caught pokemon",
			callback:    commandPokedex,
		},
		"catch": {
			name:        "catch <pokemon-name>",
			description: "try and catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <name>",
			description: "inspect pokemon from pokedex",
			callback:    commandInspect,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 map locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 map locations",
			callback:    commandMapb,
		},
		"cache": {
			name:        "cache",
			description: "prints the contents of the cache",
			callback:    commandPrintCache,
		},
		"explore": {
			name:        "explore <location-name>",
			description: "prints pokemon in area",
			callback:    commandExplore,
		},
	}
	return commands
}
