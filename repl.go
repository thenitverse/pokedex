package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedex/internal/pokeapi"
	"strings"
	"time"
)

type config struct {
	pokeapiClient pokeapi.Client
	Next          string
	Previous      string
	pokedex       map[string]pokeapi.Pokemon
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 location areas",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Explore a location area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Displays a catch message",
			callback:    commandCatch,
		},
	}
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Second*5, time.Minute*5),
		pokedex:       make(map[string]pokeapi.Pokemon),
	}
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		words := cleanInput(text)
		if len(words) == 0 {
			continue
		}
		commands := getCommands()
		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		command, ok := commands[commandName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		err := command.callback(&cfg, args...)
		if err != nil {
			fmt.Println(err)
		}

	}
}

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	words := strings.Fields(lower)
	return words

}
