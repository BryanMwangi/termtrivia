package repo

import (
	"github.com/BryanMwangi/qa/server/db"
	"github.com/BryanMwangi/qa/server/db/models"
)

func GetQuestions() models.Questions {
	return db.Questions.Get("questions").(models.Questions)
}

// we unload the questions from the cache and return them as a response without including
// the answer to avoid cheating and exposing the answer
func ResponseQuestions() []models.QuestionResponse {
	questions := GetQuestions()
	var response []models.QuestionResponse
	for _, question := range questions {
		response = append(response, models.QuestionResponse{
			ID:       question.ID,
			Question: question.Question,
			Category: question.Category,
			Options:  question.Options,
			Points:   question.Points,
		})
	}
	return response
}

func CheckAnswer(questionID int, answer int) (bool, int) {
	questions := GetQuestions()
	for _, question := range questions {
		if question.ID == questionID {
			if question.Answer == answer {
				return true, question.Points
			}
		}
	}
	return false, 0
}
