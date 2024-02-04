package pokeapi

type LocationAreaList struct {
	List
	Areas []Area `json:"results"`
}

type LocationArea struct {
	Item
	Area Area `json:"results"`
}

type Area struct {
	Item
	Pokemon_encounters []struct {
		Pokemon Pokemon `json:"pokemon"`
	} `json:"pokemon_encounters"`
}
