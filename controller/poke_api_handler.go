package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"testing_go/models"
	"testing_go/utils"

	"github.com/gorilla/mux"
)

var (
	ErrPokemonNotFound = errors.New("pokemon not found")
	ErrPokeApiFailure  = errors.New("unexpected response in PokeApi")
)

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	_, err = w.Write(response)
	if err != nil {
		log.Fatal(err)
	}
}

func GetPokemon(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	apiPokemon, err := GetPokemonFromPokeAPI(id)

	if errors.Is(err, ErrPokemonNotFound) {
		RespondWithJSON(w, http.StatusNotFound, fmt.Sprintf("Pokemon not found : %s", id))
	}

	if err != nil {
		RespondWithJSON(w, http.StatusInternalServerError, fmt.Sprintf("error while calling pokeapi: %s", err))
	}

	parsedPokemon, err := utils.ParsePokemon(apiPokemon)

	if err != nil {
		RespondWithJSON(w, http.StatusInternalServerError, fmt.Sprintf("error found: %s", err.Error()))
	}
	RespondWithJSON(w, http.StatusOK, parsedPokemon)
}

func GetPokemonFromPokeAPI(id string) (models.PokeApiPokemonResponse, error) {
	request := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", id)

	response, err := http.Get(request)

	if response.StatusCode == http.StatusNotFound {
		return models.PokeApiPokemonResponse{}, ErrPokemonNotFound
	}
	if response.StatusCode != http.StatusOK {
		return models.PokeApiPokemonResponse{}, ErrPokeApiFailure
	}
	if err != nil {
		return models.PokeApiPokemonResponse{}, err
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return models.PokeApiPokemonResponse{}, err
	}

	var apiPokemon models.PokeApiPokemonResponse

	err = json.Unmarshal(body, &apiPokemon)
	if err != nil {
		return models.PokeApiPokemonResponse{}, err
	}
	return apiPokemon, nil
}
