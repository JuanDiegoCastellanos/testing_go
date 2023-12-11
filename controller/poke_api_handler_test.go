package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"testing_go/models"

	"github.com/gorilla/mux"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/require"
)

func TestGetPokemonFromPokeApiSuccess(t *testing.T) {

	c := require.New(t)

	pokemon, err := GetPokemonFromPokeAPI("bulbasaur")

	c.NoError(err)

	body, err := os.ReadFile("samples/poke_api_readed.json")
	c.NoError(err)

	var expect models.PokeApiPokemonResponse

	err = json.Unmarshal([]byte(body), &expect)
	c.NoError(err)

	c.Equal(expect, pokemon)

}
func TestGetPokemonFromPokeApiSuccessWithMocks(t *testing.T) {
	c := require.New(t)
	// activar el mock
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	id := "bulbasaur"

	body, err := os.ReadFile("samples/pokeapi_response.json")
	c.NoError(err)

	request := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", id)

	httpmock.RegisterResponder("GET", request, httpmock.NewStringResponder(200, string(body)))

	pokemon, err := GetPokemonFromPokeAPI(id)
	c.NoError(err)

	var expected models.PokeApiPokemonResponse

	err = json.Unmarshal([]byte(body), &expected)
	c.NoError(err)

	c.Equal(expected, pokemon)

}

func TestGetPokemonFromPokeApiInternalServerError(t *testing.T) {
	c := require.New(t)
	// activar el mock
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	id := "bulbasaur"

	body, err := os.ReadFile("samples/pokeapi_response.json")
	c.NoError(err)

	request := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", id)

	httpmock.RegisterResponder("GET", request, httpmock.NewStringResponder(500, string(body)))

	_, err = GetPokemonFromPokeAPI(id)
	c.NotNil(err)
	c.EqualError(ErrPokeApiFailure, err.Error())

}

func TestGetPokemonFromPokeApiPokemonNotFound(t *testing.T) {
	c := require.New(t)
	// activar el mock
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	id := "bulbasau"

	body, err := os.ReadFile("samples/pokeapi_response.json")
	c.NoError(err)

	request := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", id)

	httpmock.RegisterResponder("GET", request, httpmock.NewStringResponder(404, string(body)))

	_, err = GetPokemonFromPokeAPI(id)
	c.NotNil(err)
	c.EqualError(ErrPokemonNotFound, err.Error())

}

func TestGetPokemon(t *testing.T) {
	c := require.New(t)

	r, err := http.NewRequest("GET", "/pokemon/{id}", nil)
	c.NoError(err)

	w := httptest.NewRecorder()

	vars := map[string]string{
		"id": "bulbasaur",
	}
	r = mux.SetURLVars(r, vars)

	GetPokemon(w, r)

	expectedBodyResponse, err := os.ReadFile("samples/api_response.json")
	c.NoError(err)

	var expectedPokemon models.Pokemon

	err = json.Unmarshal([]byte(expectedBodyResponse), &expectedPokemon)
	c.NoError(err)

	var actualPokemon models.Pokemon

	err = json.Unmarshal([]byte(w.Body.Bytes()), &actualPokemon)
	c.NoError(err)

	c.Equal(http.StatusOK, w.Code)
	c.Equal(expectedPokemon, actualPokemon)
}
