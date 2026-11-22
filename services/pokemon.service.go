package services

import (
	"errors"
	"strconv"

	"pokeapi/api"
	"pokeapi/model"
)


func SearchPokemon(name string) ([]model.Pokemon, error) {
	pokemon, err := GetPokemon(name)
	
	if errors.Is(err, model.ErrNotFound) && name != "" {
		return []model.Pokemon{}, nil
	}
	if err != nil {
		return nil, err
	}
	return pokemon, nil
}

// GetCharacter vérifie l'identifiant reçu dans l'URL, puis appelle l'API.
func GetPokemon(idParam string) (model.Pokemon, error) {
	id, err := strconv.Atoi(idParam)
	if err != nil || id < 1 {
		return model.Pokemon{}, model.ErrBadRequest
	}
	return api.GetPokemon(id)
}
