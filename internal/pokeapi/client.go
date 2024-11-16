package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/bdibon/pokedex/internal/pokecache"
)

const baseUrl = "https://pokeapi.co/api/v2/location-area/"

var cache = pokecache.NewCache(5 * time.Second)

type LocationArea result

type result struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type locationAreaResponse struct {
	Count    int            `json:"count"`
	Next     string         `json:"next"`
	Previous string         `json:"previous"`
	Results  []LocationArea `json:"results"`
}

type Pokemon struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type PokemonEncounter struct {
	Pokemon Pokemon `json:"pokemon"`
}

type locationAreaDetailsResponse struct {
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

func GetLocationAreas(offset, limit int) ([]LocationArea, error) {
	url := baseUrl + fmt.Sprintf("?offset=%d&limit=%d", offset, limit)
	data, err := getDataFromAPI(url)
	if err != nil {
		return nil, fmt.Errorf("error retrieving data: %w", err)
	}

	var lar locationAreaResponse
	err = json.Unmarshal(data, &lar)
	if err != nil {
		return nil, errors.New("error decoding response")
	}
	return lar.Results, nil
}

func GetPokemonEncounters(area string) ([]PokemonEncounter, error) {
	url := "https://pokeapi.co/api/v2/location-area/" + area
	data, err := getDataFromAPI(url)
	if err != nil {
		return nil, fmt.Errorf("error retrieving data: %w", err)
	}

	var ladr locationAreaDetailsResponse
	err = json.Unmarshal(data, &ladr)
	if err != nil {
		return nil, errors.New("error decoding response")
	}
	return ladr.PokemonEncounters, nil
}

func getDataFromAPI(url string) ([]byte, error) {
	cached, ok := cache.Get(url)
	if ok {
		return cached, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return nil, errors.New("network error")
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		if err != io.EOF {
			return nil, errors.New("error reading response body")
		}
	}
	cache.Add(url, data)
	return data, nil
}
