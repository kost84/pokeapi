package services
import (
	"strconv"
	"strings"
	"pokeapi/api"
	"pokeapi/model"
)
func SearchPokemon(name string) ([]model.Pokemon, error) {

	tousPokemon, err := api.GetAllPokemon()
	if err != nil {
		return nil, err
	}
	name = strings.ToLower(strings.TrimSpace(name))
	if name == "" {
		return tousPokemon, nil
	}
	var results []model.Pokemon
	for _, p := range tousPokemon {
		if strings.Contains(strings.ToLower(p.Name.Fr), name) {
			results = append(results, p)
		}
	}
	return results, nil
}

// GetCharacter vérifie l'identifiant reçu dans l'URL, puis appelle l'API.
func GetPokemon(idParam string) (model.Pokemon, error) {
	id, err := strconv.Atoi(idParam)
	if err != nil || id < 1 {
		return model.Pokemon{}, model.ErrBadRequest
	}
	return api.GetPokemon(id)
}