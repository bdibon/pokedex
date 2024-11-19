package pokeapi

type Status int

const (
	ResourceNotFound Status = iota + 1
)

type PokeApiError struct {
	message string
	Code    Status
}

func (pae PokeApiError) Error() string {
	return pae.message
}
