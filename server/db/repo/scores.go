package repo

import (
	"sort"

	"github.com/BryanMwangi/qa/server/db"
	"github.com/BryanMwangi/qa/server/db/models"
)

func GetScores() models.Scores {
	return db.Scores.Get("scores").(models.Scores)
}

func AddScore(user models.User, count int) {
	scores := GetScores()
	userFound := false

	// Check if the user already exists in the scores array
	for i, score := range scores.Scores {
		if score.UserName == user.UserName {
			scores.Scores[i].Score += count
			userFound = true

			if scores.Scores[i].Score > scores.Top {
				scores.Top = scores.Scores[i].Score
			}

			if scores.Bottom == 0 {
				scores.Bottom = scores.Scores[i].Score
			} else if scores.Scores[i].Score < scores.Bottom {
				scores.Bottom = scores.Scores[i].Score
			}
			break
		}
	}

	// If the user was not found, add a new score entry
	if !userFound {
		scores.Scores = append(scores.Scores, user)
		// Update the top and bottom scores for the new entry
		if count > scores.Top {
			scores.Top = count
		}
		if scores.Bottom == 0 || count < scores.Bottom {
			scores.Bottom = count
		}
	}

	scores.Average = (scores.Top + scores.Bottom) / 2

	db.Scores.Set("scores", scores)
}

// Get the score of a user
func GetScore(userName string) models.User {
	scores := GetScores()
	for _, score := range scores.Scores {
		if score.UserName == userName {
			return score
		}
	}
	return models.User{}
}

// Will fetch all top players sorted by score
func GetTopScorers() []models.UserResponse {
	scores := GetScores()
	if len(scores.Scores) == 0 {
		return []models.UserResponse{}
	}
	var topScorers []models.UserResponse
	for _, score := range scores.Scores {
		user := GetUser(score.UserName)
		// We don't want to expose the uuid of the user
		// we construct a response with only the username and score
		resp := models.UserResponse{
			UserName: user.UserName,
			Score:    score.Score,
		}
		topScorers = append(topScorers, resp)
	}
	// sort the top scorers by score
	sort.Slice(topScorers, func(i, j int) bool {
		return topScorers[i].Score > topScorers[j].Score
	})
	return topScorers
}

func GetGameStats() models.Stats {
	scores := GetScores()
	return models.Stats{
		Average: scores.Average,
		Bottom:  scores.Bottom,
		Top:     scores.Top,
	}
}
