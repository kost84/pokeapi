package controllers

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"pokeapi/helpers"
	"pokeapi/model"
	"pokeapi/services"
)

// ListPage : données de la page de liste.
type ListPage struct {
	Search     string // recherche saisie, pour la réafficher dans le formulaire
	Pokemon []model.Pokemon
}

// ListPokemon affiche la liste des personnages, filtrée par ?nom=… si besoin.
func ListPokemon(w http.ResponseWriter, r *http.Request) {
	search := strings.TrimSpace(r.URL.Query().Get("nom"))
	pokemon, err := services.SearchPokemon(search)
	if err != nil {
		renderAPIError(w, err)
		return
	}
	helpers.Render(w, http.StatusOK, "list", ListPage{Search: search, Pokemon: pokemon})
}

// ShowPokemon affiche la fiche du personnage dont l'identifiant est dans l'URL.
func ShowPokemon(w http.ResponseWriter, r *http.Request) {
	pokemon, err := services.GetPokemon(r.PathValue("id"))
	if err != nil {
		renderAPIError(w, err)
		return
	}
	helpers.Render(w, http.StatusOK, "character", pokemon)
}

// NotFound affiche la page 404 pour toute URL inconnue du site.
func NotFound(w http.ResponseWriter, r *http.Request) {
	helpers.RenderError(w, http.StatusNotFound, "Cette page n'existe pas.")
}

// renderAPIError traduit une erreur connue en page d'erreur, sur la route même.
func renderAPIError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, model.ErrNotFound):
		helpers.RenderError(w, http.StatusNotFound,
			"Ce personnage n'existe pas.")
	case errors.Is(err, model.ErrBadRequest):
		helpers.RenderError(w, http.StatusBadRequest,
			"Identifiant de personnage invalide.")
	case errors.Is(err, model.ErrUnavailable):
		log.Println("The Rick and Morty API :", err)
		helpers.RenderError(w, http.StatusServiceUnavailable,
			"Service momentanément indisponible.")
	default:
		log.Println("Erreur inattendue :", err)
		helpers.RenderError(w, http.StatusInternalServerError,
			"Une erreur est survenue.")
	}
}
