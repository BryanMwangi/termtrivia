package models

// I am aware of the limitations of this approach but this will have to work for now
// If we need to have high availability, perhaps a binary search tree would be better
// but for the sake of a demo, I will use a simple array
type Scores struct {
	Scores  []User `json:"scores"`
	Average int    `json:"average"`
	Top     int    `json:"top"`
	Bottom  int    `json:"bottom"`
}

type Stats struct {
	Average int `json:"average"`
	Top     int `json:"top"`
	Bottom  int `json:"bottom"`
}
