package models

type Question struct {
	// The number of the question
	ID int `json:"id"`
	// The question itself
	Question string `json:"question"`
	// The category of the question
	Category string `json:"category"`
	// options are an array of objects that specify the different answer options for
	// the question
	Options []Option `json:"options"`
	// we use int to specify the answer to reduce the memory footprint
	Answer int `json:"answer"`
	// Points awarded for each correct answer
	Points int `json:"points"`
}

// The questions sent to the client should not contain the answer
type QuestionResponse struct {
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

type Questions []Question
