package handlers

import (
	"net/http"

	"github.com/BryanMwangi/pine"
	"github.com/BryanMwangi/qa/server/db/repo"
	"github.com/BryanMwangi/qa/server/utils"
)

func hello(c *pine.Ctx) error {
	return c.JSON("Hello, world!")
}

func getUser(c *pine.Ctx) error {
	userName := c.Params("username")
	if userName == "" {
		return c.JSON("Username is required", 400)
	}
	pass, name, err := utils.ValidateName(userName)
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}
	if !pass {
		return c.JSON("Invalid username", 400)
	}
	user := repo.CreateOrFetchUser(name)
	return c.JSON(user)
}

func getQuestions(c *pine.Ctx) error {
	questions := repo.ResponseQuestions()
	return c.JSON(questions)
}

func answerQuestion(c *pine.Ctx) error {
	questionID, err := c.ParamsInt("id")
	if err != nil {
		return c.JSON("Question ID is required", 400)
	}
	answer, err := c.ParamsInt("answer")
	if err != nil {
		return c.JSON("Answer is required", 400)
	}
	correct, points := repo.CheckAnswer(questionID, answer)
	if correct {
		userName := c.Locals("userName")
		if userName == nil {
			return c.SendStatus(401)
		}
		user := repo.GetUser(userName.(string))
		repo.AddScore(user, points)
	}
	return c.JSON(correct)
}

func getScore(c *pine.Ctx) error {
	userName := c.Locals("userName")
	if userName == nil {
		return c.SendStatus(401)
	}
	score := repo.GetScore(userName.(string))
	return c.JSON(score)
}

func getTopScorers(c *pine.Ctx) error {
	topScorers := repo.GetTopScorers()
	return c.JSON(topScorers)
}

func getGameStats(c *pine.Ctx) error {
	stats := repo.GetGameStats()
	return c.JSON(stats)
}
