package handlers

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/BryanMwangi/qa/cliapp/client"
	"github.com/BryanMwangi/qa/cliapp/client/models"
	"github.com/BryanMwangi/qa/cliapp/utils"
)

var (
	currentQuestion int
	questionPool    models.Questions
	GameComplete    bool
	TotalQuestions  int
	CorrectAnswers  int
)

func FetchQuestions() error {
	// reset game state
	GameComplete = false

	questionsUri := client.GenerateUri("/questions")
	client.Client.Request().SetRequestURI(questionsUri)

	//fetch questions
	if err := client.Client.SendRequest(); err != nil {
		return err
	}
	utils.ShowLoader("Loading...")
	_, body, err := client.Client.ReadResponse()
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &questionPool)
	if err != nil {
		return err
	}
	currentQuestion = questionPool[0].ID
	TotalQuestions = len(questionPool)
	CorrectAnswers = 0
	utils.StopLoader(200 * time.Millisecond)
	return nil
}

// GetNextQuestion returns the next question in the question pool
func GetNextQuestion() models.Question {
	return questionPool[currentQuestion]
}

func AnswerQuestion(questionID int, answer int) bool {
	// increase current question count
	currentQuestion++

	// check if game is complete
	if currentQuestion >= len(questionPool) {
		// set game complete
		GameComplete = true
		// reset current question count
		currentQuestion = 0
	}
	answerString := fmt.Sprintf("/answer/%d/%d", questionID, answer)
	answerUrl := client.GenerateUri(answerString)
	client.Client.Request().SetRequestURI(answerUrl)
	if err := client.Client.SendRequest(); err != nil {
		return false
	}
	_, body, err := client.Client.ReadResponse()
	if err != nil {
		fmt.Println("Error reading response", err)
		return false
	}
	var answerResponse bool
	err = json.Unmarshal(body, &answerResponse)
	if err != nil {
		fmt.Println("Error unmarshalling response", err)
		return false
	}
	if answerResponse {
		CorrectAnswers++
	}
	return answerResponse
}
