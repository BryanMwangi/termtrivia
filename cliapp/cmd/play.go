package cmd

import (
	"fmt"
	"log"

	"github.com/BryanMwangi/pine/logger"
	"github.com/BryanMwangi/qa/cliapp/client/handlers"
	"github.com/BryanMwangi/qa/cliapp/client/models"
	"github.com/manifoldco/promptui"
)

func playGame() {
	for {
		// Fetch a question and its options
		question := handlers.GetNextQuestion()

		// Display the question
		fmt.Println("\nQuestion:", question.Question)

		// Use promptui to display options and capture the user's choice
		prompt := promptui.Select{
			Label: "Choose your answer",
			Items: extractOptions(question.Options),
		}

		pos, _, err := prompt.Run()
		if err != nil {
			log.Fatalf("Prompt failed: %v\n", err)
		}
		res := handlers.AnswerQuestion(question.ID, pos)
		// Display the chosen answer
		isCorrect(res)

		if handlers.GameComplete {
			fmt.Println("Fetching your score...")
			user, err := handlers.FetchUserScore()
			if err != nil {
				fmt.Println("Failed to fetch your score")
				fmt.Println(err)
			}
			fmt.Printf("You got %d correct answers out of %d questions\n", handlers.CorrectAnswers, handlers.TotalQuestions)
			fmt.Printf("\nYour current total score is: %d\n\n", user.Score)
			fmt.Println("Thanks for playing!")
			break
		}
	}
}

func extractOptions(options []models.Option) []string {
	var optionsStrings []string
	for _, option := range options {
		optionsStrings = append(optionsStrings, option.Option)
	}
	return optionsStrings
}

func isCorrect(result bool) {
	if result {
		logger.Success("Correct!")
		return
	}
	logger.Error("Incorrect!")
}
