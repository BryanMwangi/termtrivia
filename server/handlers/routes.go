package handlers

import (
	"github.com/BryanMwangi/pine"
)

func Routes(app *pine.Server) {
	app.Get("/hello", hello)
	app.Get("/user/:username", getUser)
	app.Get("/questions", getQuestions)
	app.Get("/topScorers", getTopScorers)
	app.Get("/gameStats", getGameStats)
	app.Get("/answer/:id/:answer", authenticateRequest(), answerQuestion)
	app.Get("/score", authenticateRequest(), getScore)
}
