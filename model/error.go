package model

import (
	"errors"
)

var (
 ErrNotFound = errors.New("Erreur 404 . page introuvable")
 ErrBadRequest = errors.New("Erreur de requete")
 ErrTimeout = errors.New("délai d'attente dépassé")
 ErrUnavailable= errors.New("service indisponible")

)
