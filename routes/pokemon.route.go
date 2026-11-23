package routes

import (
	"net/http"
	"pokeapi/controllers"
)

func characterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /pokemon", controllers.ListPokemon)
	mux.HandleFunc("GET /pokemon/{id}", controllers.ShowPokemon)
}