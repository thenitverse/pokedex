package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("not enough args.")
	}
	name := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	pokemonData, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}
	roll := rand.Intn(100)
	catchChance := 100 - (pokemonData.BaseExperience / 2)
	if catchChance < 10 {
		catchChance = 10
	}
	if roll < catchChance {
		cfg.pokedex[name] = pokemonData
		fmt.Printf("%s was caught!\n", name)
	} else {
		fmt.Printf("%s escaped!\n", name)
	}
	return nil

}
