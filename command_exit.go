package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"os"
)

type locationAreaResponse struct {
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
	} `json:"results"`
}

func commandExit(cfg *config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil

}
func commandMap(cfg *config, args ...string) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if cfg.Next != "" {
		url = cfg.Next
	}
	return fetchLocations(cfg, url)

}
func fetchLocations(cfg *config, url string) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	var data locationAreaResponse
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&data)
	if err != nil {
		return err
	}
	if data.Next != nil {
		cfg.Next = *data.Next
	} else {
		cfg.Next = ""
	}
	if data.Previous != nil {
		cfg.Previous = *data.Previous
	} else {
		cfg.Previous = ""
	}
	for _, item := range data.Results {
		fmt.Println(item.Name)
	}
	return nil
}

func commandMapb(cfg *config, args ...string) error {
	if cfg.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	return fetchLocations(cfg, cfg.Previous)
}
func commandHelp(cfg *config, args ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println()
	fmt.Println("Usage:")
	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)

	}
	return nil
}
