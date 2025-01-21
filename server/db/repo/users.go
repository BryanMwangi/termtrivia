package repo

import (
	"fmt"

	"github.com/BryanMwangi/qa/server/db"
	"github.com/BryanMwangi/qa/server/db/models"
	"github.com/google/uuid"
)

func CreateOrFetchUser(userName string) models.User {
	// first check if the user exists in the cache
	if user := db.Users.Get(userName); user != nil {
		return user.(models.User)
	}
	user := models.User{
		ID:       uuid.New(),
		UserName: userName,
	}
	// we can also add the user's score count and it will always start with 0
	AddScore(user, 0)
	db.Users.Set(user.UserName, user)
	return user
}

func GetUser(userName string) models.User {
	if user := db.Users.Get(userName); user != nil {
		return user.(models.User)
	}
	fmt.Println("User not found")
	return models.User{}
}
