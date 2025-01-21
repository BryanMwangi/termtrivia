package models

type User struct {
	ID       string `json:"id"`
	UserName string `json:"username"`
	Score    int    `json:"score"`
}

type Stats struct {
	Average int `json:"average"`
	Top     int `json:"top"`
	Bottom  int `json:"bottom"`
}

type LeaderboardUser struct {
	UserName string `json:"username"`
	Score    int    `json:"score"`
}
