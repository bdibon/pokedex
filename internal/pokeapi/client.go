package pokeapi

import (
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/bdibon/pokedex/internal/pokecache"
)

var cache = pokecache.NewCache(5 * time.Second)

func getDataFromAPI(url string) ([]byte, error) {
	cached, ok := cache.Get(url)
	if ok {
		return cached, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return nil, errors.New("network error")
	}
	if res.StatusCode != http.StatusOK {
		if res.StatusCode == http.StatusNotFound {
			return nil, PokeApiError{
				"resource not found",
				ResourceNotFound,
			}
		}
		return nil, errors.New("unsuccessful response status")
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
