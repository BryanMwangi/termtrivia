package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/BryanMwangi/qa/cliapp/client"
	"github.com/BryanMwangi/qa/cliapp/client/models"
)

func FetchUserScore() (models.User, error) {
	scoreUri := client.GenerateUri("/score")
	client.Client.Request().SetRequestURI(scoreUri)

	if err := client.Client.SendRequest(); err != nil {
		return models.User{}, err
	}
	_, body, err := client.Client.ReadResponse()
	if err != nil {
		return models.User{}, err
	}
	var user models.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func GameStats() (models.Stats, error) {
	statsUri := client.GenerateUri("/gameStats")
	client.Client.Request().SetRequestURI(statsUri)

	if err := client.Client.SendRequest(); err != nil {
		return models.Stats{}, err
	}
	_, body, err := client.Client.ReadResponse()
	if err != nil {
		return models.Stats{}, err
	}
	var stats models.Stats
	err = json.Unmarshal(body, &stats)
	if err != nil {
		return models.Stats{}, err
	}

	return stats, nil
}

func UserStats() error {
	score, err := FetchUserScore()
	if err != nil {
		return err
	}
	stats, err := GameStats()
	if err != nil {
		return err
	}
	fmt.Printf("\nYour score is: %d\n\n", score.Score)
	fmt.Println(compareUserAgainstAverage(score.Score, stats.Average))
	if score.Score != 0 && score.Score == stats.Top {
		fmt.Println("You are on top of the leaderboard!")
	}

	return nil
}

func compareUserAgainstAverage(score int, average int) string {
	if score == 0 || average == 0 {
		return "You have not played yet"
	}

	var diff int
	var above bool
	if score > average {
		above = true
		diff = score - average
	} else {
		above = false
		diff = average - score
	}

	// calculate the player's average score percentage
	perc := (float32(diff) / float32(average)) * 100

	if above {
		return fmt.Sprintf("Your score is %.2f%% higher than the average", perc)
	}
	return fmt.Sprintf("Your score is %.2f%% less than the average. You need to work harder!", perc)
}
