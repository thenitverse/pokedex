package pokeapi

import (
	"encoding/json"
	//"fmt"
	"io"
	"net/http"
)

type LocationArea struct {
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}
type PokemonEncounter struct {
	Pokemon LocationPokemon `json:"pokemon"`
}
type LocationPokemon struct {
	Name string `json:"name"`
}
type Pokemon struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
}

func (c *Client) GetLocationArea(areaName string) (LocationArea, error) {
	url := baseURL + "/location-area/" + areaName
	if data, ok := c.cache.Get(url); ok {
		var locationArea LocationArea

		err := json.Unmarshal(data, &locationArea)
		if err != nil {
			return LocationArea{}, err
		}
		return locationArea, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationArea{}, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationArea{}, err
	}
	c.cache.Add(url, data)
	var locationArea LocationArea
	err = json.Unmarshal(data, &locationArea)
	if err != nil {
		return LocationArea{}, err
	}
	return locationArea, nil

}
func (c *Client) GetPokemon(name string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + name
	if data, ok := c.cache.Get(url); ok {
		var PokemonName Pokemon
		err := json.Unmarshal(data, &PokemonName)
		if err != nil {
			return Pokemon{}, err
		}
		return PokemonName, nil

	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}
	c.cache.Add(url, data)
	var PokemonName Pokemon
	err = json.Unmarshal(data, &PokemonName)
	if err != nil {
		return Pokemon{}, err
	}
	return PokemonName, nil

}
