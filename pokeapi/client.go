package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

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
	url := "https://pokeapi.co/api/v2/location-area/" + fmt.Sprintf("?offset=%d&limit=%d", offset, limit)
	res, err := http.Get(url)
	if err != nil {
		return nil, errors.New("getLocationAreas: network error")
	}

	var data locationAreaResponse
	dec := json.NewDecoder(res.Body)
	if err = dec.Decode(&data); err != nil {
		return nil, errors.New("getLocationAreas: error decoding response")
	}
	return data.Results, nil
}
