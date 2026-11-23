# Pokédex en Go (Tyradex API)

Ce projet est une application web développée en Go qui consomme l'API REST Tyradex afin de proposer un Pokédex en ligne en français. L'application permet de consulter la liste des Pokémon, de visualiser la fiche détaillée de chacun, d'effectuer une recherche par nom et de gérer les erreurs HTTP de manière explicite.

## Origine et inspiration du code

L'architecture et la structure de ce projet sont directement inspirées du code source de référence fourni dans l'archive "Guide pratique - Code source - Rick & Morty" (`Guide_pratique_-_Code_source__-_Rick__Morty_(1).zip`), issu de la capsule de cours "Les APIs & Go".

Le découpage technique reprend les principes démontrés dans ce guide pratique :
- Séparation stricte des responsabilités en couches (API, Modèles, Services, Contrôleurs, Routes, Helpers).
- Initialisation unique d'un client HTTP avec timeout et configuration des en-têtes (User-Agent).
- Centralisation des erreurs applicatives avec la bibliothèque standard (errors.Is) et rendu d'une page d'erreur dédiée sans redirection.
- Rendu serveur via html/template avec chargement unique des gabarits au démarrage.

Cette base a été adaptée pour répondre aux spécificités de l'API Tyradex (champs utiles des Pokémon, filtrage en mémoire pour la recherche, gestion des statuts de réponse).

## Architecture en couches

Le projet respecte l'arborescence suivante :

```text
pokedex/
├── main.go                       Point d'entrée du serveur web
├── api/
│   └── tyradex.api.go            Appels HTTP vers l'API Tyradex (client HTTP, décodage JSON)
├── model/ (ou models/)
│   ├── pokeinfo.go               Structures de données (Pokemon, Name, Sprites, Type)
│   └── error.go                  Déclaration des erreurs sentinelles (ErrNotFound, ErrBadRequest, etc.)
├── services/
│   └── pokemon.service.go        Logique métier (validation de l'ID, filtrage de recherche)
├── controllers/
│   └── pokemon.controller.go     Traitement des requêtes HTTP et sélection des vues/statuts
├── routes/
│   ├── main.route.go             Routeur principal et service des fichiers statiques
│   └── pokemon.route.go          Définition des routes de l'application
├── helpers/
│   └── pokemon.helper.go         Chargement et exécution des templates HTML
├── templates/
│   ├── list.html                 Page listant les Pokémon et formulaire de recherche
│   ├── pokemon.html              Page de détail d'un Pokémon
│   └── error.html                Page d'erreur générique
└── asset/css/ (ou assets/css/)
    └── style.css                 Feuille de styles CSS
```

## Prérequis et installation

- Go 1.22 ou supérieur installé sur la machine.
- Uniquement la bibliothèque standard de Go (aucun framework ou bibliothèque tierce externe).

## Lancement de l'application

1. Cloner ou ouvrir le dossier du projet dans un terminal.
2. Lancer le serveur avec la commande suivante :

```bash
go run .
```

3. Le serveur démarre sur le port 8080 : `http://localhost:8080`

## Exemples d'URL et fonctionnalités

- Redirection d'accueil :
  `http://localhost:8080/`
  Redirige de manière permanente (code HTTP 301) vers `/pokemon`.

- Liste des Pokémon :
  `http://localhost:8080/pokemon`
  Affiche l'ensemble des Pokémon disponibles avec leur numéro, nom français et visuel.

- Fiche d'un Pokémon existant :
  `http://localhost:8080/pokemon/25`
  Affiche la fiche détaillée de Pikachu (code HTTP 200) avec sa catégorie, son image et ses types.

- Recherche par nom :
  `http://localhost:8080/pokemon?nom=pika`
  Filtre la liste en conservant les Pokémon dont le nom français contient le terme recherché (insensible à la casse et sans espaces superflus).

- Erreur - Identifiant invalide :
  `http://localhost:8080/pokemon/abc`
  Renvoie un code HTTP 400 avec le message "Numéro de Pokémon invalide." sans interroger l'API Tyradex.

- Erreur - Pokémon inexistant :
  `http://localhost:8080/pokemon/9999`
  Renvoie un code HTTP 404 avec le message "Ce Pokémon n'existe pas.".

- Erreur - Route inconnue :
  `http://localhost:8080/page-inconnue`
  Renvoie un code HTTP 404 avec le message "Cette page n'existe pas.".

## Source des données

Les données et visuels utilisés proviennent de l'API Tyradex (https://tyradex.app). Tyradex est mentionné comme source des données conformément aux conditions d'utilisation.
