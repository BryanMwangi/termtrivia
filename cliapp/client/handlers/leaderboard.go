package handlers

import (
	"encoding/json"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/BryanMwangi/qa/cliapp/client"
	"github.com/BryanMwangi/qa/cliapp/client/models"
)

func FetchLeaderboard() ([]models.LeaderboardUser, error) {
	leaderboardUri := client.GenerateUri("/topScorers")
	client.Client.Request().SetRequestURI(leaderboardUri)

	if err := client.Client.SendRequest(); err != nil {
		return []models.LeaderboardUser{}, err
	}
	_, body, err := client.Client.ReadResponse()
	if err != nil {
		return []models.LeaderboardUser{}, err
	}
	var leaderboard []models.LeaderboardUser
	err = json.Unmarshal(body, &leaderboard)
	if err != nil {
		return []models.LeaderboardUser{}, err
	}
	return leaderboard, nil
}

func GetLeaderboard() error {
	leaderBoard, err := FetchLeaderboard()
	if err != nil {
		return err
	}

	fmt.Printf(`
	▗▖   ▗▄▄▄▖ ▗▄▖ ▗▄▄▄ ▗▄▄▄▖▗▄▄▖ ▗▄▄▖  ▗▄▖  ▗▄▖ ▗▄▄▖ ▗▄▄▄ 
	▐▌   ▐▌   ▐▌ ▐▌▐▌  █▐▌   ▐▌ ▐▌▐▌ ▐▌▐▌ ▐▌▐▌ ▐▌▐▌ ▐▌▐▌  █
	▐▌   ▐▛▀▀▘▐▛▀▜▌▐▌  █▐▛▀▀▘▐▛▀▚▖▐▛▀▚▖▐▌ ▐▌▐▛▀▜▌▐▛▀▚▖▐▌  █
	▐▙▄▄▖▐▙▄▄▖▐▌ ▐▌▐▙▄▄▀▐▙▄▄▖▐▌ ▐▌▐▙▄▞▘▝▚▄▞▘▐▌ ▐▌▐▌ ▐▌▐▙▄▄▀
	
`)

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "Rank\tUsername\tScore\t")
	fmt.Fprintln(w, "----\t--------\t-----\t")
	for i, user := range leaderBoard {
		fmt.Fprintf(w, "| %d \t| %s \t| %d \t|\n", i+1, user.UserName, user.Score)
	}

	w.Flush()
	return nil
}
