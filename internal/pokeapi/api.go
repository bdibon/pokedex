package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
)

const baseUrl = "https://pokeapi.co/api/v2/location-area/"

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

type PokemonEncounter struct {
	Pokemon struct {
		Name string `json:"name"`
		Url  string `json:"url,omitempty"`
	} `json:"pokemon"`
}

type locationAreaDetailsResponse struct {
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

type Pokemon struct {
	Abilities []struct {
		Ability struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"ability"`
		IsHidden bool `json:"is_hidden"`
		Slot     int  `json:"slot"`
	} `json:"abilities"`
	BaseExperience int `json:"base_experience"`
	Cries          struct {
		Latest string `json:"latest"`
		Legacy string `json:"legacy"`
	} `json:"cries"`
	Forms []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"forms"`
	GameIndices []struct {
		GameIndex int `json:"game_index"`
		Version   struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"version"`
	} `json:"game_indices"`
	Height    int `json:"height"`
	HeldItems []struct {
		Item struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"item"`
		VersionDetails []struct {
			Rarity  int `json:"rarity"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"held_items"`
	ID                     int    `json:"id"`
	IsDefault              bool   `json:"is_default"`
	LocationAreaEncounters string `json:"location_area_encounters"`
	Moves                  []struct {
		Move struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"move"`
		VersionGroupDetails []struct {
			LevelLearnedAt  int `json:"level_learned_at"`
			MoveLearnMethod struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"move_learn_method"`
			VersionGroup struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version_group"`
		} `json:"version_group_details"`
	} `json:"moves"`
	Name          string `json:"name"`
	Order         int    `json:"order"`
	PastAbilities []any  `json:"past_abilities"`
	PastTypes     []any  `json:"past_types"`
	Species       struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"species"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	Weight int `json:"weight"`
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

func GetPokemon(name string) (Pokemon, error) {
	url := "https://pokeapi.co/api/v2/pokemon/" + name
	data, err := getDataFromAPI(url)
	if err != nil {
		return Pokemon{}, fmt.Errorf("error retrieving data: %w", err)
	}

	var pkmn Pokemon
	err = json.Unmarshal(data, &pkmn)
	if err != nil {
		return Pokemon{}, errors.New("error decoding response")
	}
	return pkmn, nil
}
