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

func GetLocationAreas(offset, limit int) ([]LocationArea, error) {
	data, err := getLocationAreasData(offset, limit)
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

func getLocationAreasData(offset, limit int) ([]byte, error) {
	ckey := fmt.Sprintf("location-area-%d-%d", offset, limit)
	cached, ok := cache.Get(ckey)
	if ok {
		return cached, nil
	}

	url := "https://pokeapi.co/api/v2/location-area/" + fmt.Sprintf("?offset=%d&limit=%d", offset, limit)
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
	cache.Add(ckey, data)
	return data, nil
}
