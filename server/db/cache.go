package db

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/BryanMwangi/pine/cache"
	"github.com/BryanMwangi/pine/logger"
	"github.com/BryanMwangi/qa/server/db/models"
)

var (
	// Pine provides a versatile built in cache system that has read and write capabilities
	// with efficient time complexities of O(1) and 0(n) for massive objects of several
	// degrees of data size for both read and write operations accross
	// multiple goroutines.
	//
	// There was no need for me to implement my own cache system given I already wrote one
	// in Pine and it would just look the same
	//
	// you can read more about the cache system here: https://github.com/BryanMwangi/pine/blob/main/cache/cache.go
	Questions *cache.Cache
	Users     *cache.Cache
	// We could further distribute the cache if needed to avoid a whole cache set and get
	// however, realistically, this is for a demo purpose but I am just saying
	//
	// Distribution will allow for fast access to the data
	Scores *cache.Cache

	//reset the cache every hour
	resetTime = 1 * time.Hour

	// next reset time
	NextReset = time.Now().Add(resetTime)
)

func Init() {
	// create a new cache with an expiration time of 1 hour
	//
	// each Q&A result will be cached for at least 1 hour by the server. So will each session
	Users = cache.New(time.Hour)
	Scores = cache.New(time.Hour)
	Questions = cache.New(time.Hour)
	// load the questions from the file
	loadQuestions()
	// load the scores to initialize the cache
	loadScores()
	// start a goroutine to reset the cache every hour
	go reset()
}

func reset() {
	ticker := time.NewTicker(resetTime)
	defer ticker.Stop()
	for {
		<-ticker.C
		Init()
		NextReset = time.Now().Add(resetTime)
		fmt.Println("Cache reset")
		fmt.Println("Next reset at: " + NextReset.String())
	}
}

func loadQuestions() {
	questionsFile, err := os.ReadFile("db/questions.json")
	if err != nil {
		logger.Error("Error loading questions: " + err.Error())
		panic(err)
	}
	var questions models.Questions
	if err := json.Unmarshal(questionsFile, &questions); err != nil {
		logger.Error("Error unmarshalling questions: " + err.Error())
		panic(err)
	}
	Questions.Set("questions", questions)
}

func loadScores() {
	var scores models.Scores
	Scores.Set("scores", scores)
}
