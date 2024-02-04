module github.com/gordjw/pokedexcli

go 1.21.6

replace pokeapi => ./internal/pokeapi

replace pokecache => ./internal/pokecache

require pokeapi v0.0.0-00010101000000-000000000000

require pokecache v0.0.0-00010101000000-000000000000 // indirect
