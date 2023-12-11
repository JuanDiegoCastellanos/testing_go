package utils

import (
	"encoding/json"
	"io/ioutil"
	"testing"
	"testing_go/models"

	"github.com/stretchr/testify/require"
)

// comando para correr test especificos
// go test ./utils/ -run=TestParserPokemonSuccess
func TestParserPokemonSuccess(t *testing.T) {
	c := require.New(t)

	body, err := ioutil.ReadFile("samples/pokeapi_response.json")
	c.NoError(err)

	var response models.PokeApiPokemonResponse

	err = json.Unmarshal([]byte(body), &response)
	c.NoError(err)

	parsedPokemon, err := ParsePokemon(response)
	c.NoError(err)

	body, err = ioutil.ReadFile("samples/api_response.json")

	c.NoError(err)

	var expected models.Pokemon
	err = json.Unmarshal([]byte(body), &expected)

	c.NoError(err)

	c.Equal(expected, parsedPokemon)
}

func TestParserPokemonTypeNotFound(t *testing.T) {
	c := require.New(t)
	body, err := ioutil.ReadFile("samples/pokeapi_response.json")
	c.NoError(err)

	var response models.PokeApiPokemonResponse

	err = json.Unmarshal([]byte(body), &response)
	c.NoError(err)

	response.PokemonType = []models.PokemonType{}

	_, err = ParsePokemon(response)
	c.NotNil(err)
	c.EqualError(ErrNotFoundPokemonType, err.Error())

}

// pruebas de rendimiento
func BenchmarkParser(b *testing.B) {
	c := require.New(b)

	body, err := ioutil.ReadFile("samples/pokeapi_response.json")
	c.NoError(err)

	var response models.PokeApiPokemonResponse

	err = json.Unmarshal([]byte(body), &response)
	c.NoError(err)

	for n := 0; n < b.N; n++ {
		_, err := ParsePokemon(response)
		c.NoError(err)

	}

}
