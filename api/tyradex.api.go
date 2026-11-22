package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"pokeapi/model" 
)

const baseURL = "https://tyradex.app/api/v1"

var client = &http.Client{Timeout: 5 * time.Second}

func GetPokemon(id int) (model.Pokemon, error) {
	endpoint := fmt.Sprintf("%s/pokemon/%d", baseURL, id)

	request, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return model.Pokemon{}, err
	}
	request.Header.Set("User-Agent", "DataDex/1.0")

	response, err := client.Do(request)
	if err != nil {
		return model.Pokemon{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return model.Pokemon{}, fmt.Errorf("réponse inattendue : %s", response.Status)
	}

	var pokemon model.Pokemon
	err = json.NewDecoder(response.Body).Decode(&pokemon)
	return pokemon, err
}

func GetAllPokemon()([]model.Pokemon,error){
	endpoint := fmt.Sprintf("%s/pokemon/", baseURL)

	request, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return []model.Pokemon{}, err
	}
	request.Header.Set("User-Agent", "DataDex/1.0")

	response, err := client.Do(request)
	if err != nil{
		return []model.Pokemon{},err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return []model.Pokemon{}, fmt.Errorf("Réponse inattendue : %s", response.Status)
	}
	var pokemon []model.Pokemon
	err = json.NewDecoder(response.Body).Decode(&pokemon)
	return pokemon, err
}