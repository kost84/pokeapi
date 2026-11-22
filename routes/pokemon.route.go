package routes

import (
	"net/http"

	"pokeapi/controllers"
)

func characterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /personnages", controllers.ListPokemon)
	mux.HandleFunc("GET /personnages/{id}", controllers.ShowPokemon)
}
