package models

type Questions []Question

type Question struct {
	ID       int      `json:"id"`
	Question string   `json:"question"`
	Category string   `json:"category"`
	Options  []Option `json:"options"`
	Points   int      `json:"points"`
}

type Option struct {
	ID     int    `json:"id"`
	Option string `json:"option"`
}
