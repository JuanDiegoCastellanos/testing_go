package utils

import (
	"errors"
	"testing_go/models"
)

var (
	ErrNotFoundPokemonType     = errors.New("Pokemon type array not found")
	ErrNotFoundPokemonTypeName = errors.New("Pokemon type name not found")
)

func ParsePokemon(apiPokemon models.PokeApiPokemonResponse) (models.Pokemon, error) {
	if len(apiPokemon.PokemonType) < 1 {
		return models.Pokemon{}, ErrNotFoundPokemonType
	}
	if apiPokemon.PokemonType[0].RefType.Name == "" {
		return models.Pokemon{}, ErrNotFoundPokemonTypeName
	}
	pokemonType := apiPokemon.PokemonType[0].RefType.Name

	abilitiesMap := map[string]int{}

	for _, stat := range apiPokemon.Stats {
		parsedAbilityName, ok := models.AllowedAbilities[stat.Stat.Name]
		if !ok {
			continue
		}
		abilitiesMap[parsedAbilityName] = stat.BaseStat
	}
	parsedPokemon := models.Pokemon{
		Id:        apiPokemon.Id,
		Name:      apiPokemon.Name,
		Power:     pokemonType,
		Abilities: abilitiesMap,
	}
	return parsedPokemon, nil
}
