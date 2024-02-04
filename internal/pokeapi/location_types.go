package pokeapi

type LocationList struct {
	List
	Locations []Location `json:"results"`
}

type Location struct {
	Item
	Region Region `json:"region"`
	Areas  []Area `json:"areas"`
}
