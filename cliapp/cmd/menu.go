package cmd

import (
	"fmt"

	"github.com/BryanMwangi/qa/cliapp/client/handlers"
	"github.com/manifoldco/promptui"
)

// main game loop
func gameLoop() {
	for {
		options := []string{"Start Game", "My Stats", "Top Scorers", "Quit"}
		prompt := promptui.Select{
			Label: "Choose an option",
			Items: options,
		}

		_, result, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		switch result {
		case "Start Game":
			handlers.FetchQuestions()
			playGame()
		case "My Stats":
			err := handlers.UserStats()
			if err != nil {
				fmt.Println(err)
			}
		case "Top Scorers":
			err := handlers.GetLeaderboard()
			if err != nil {
				fmt.Println(err)
			}
		case "Quit":
			return
		}
	}
}
