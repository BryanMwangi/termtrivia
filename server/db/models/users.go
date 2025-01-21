package models

import "github.com/google/uuid"

type Users []User

// no need for authentication at this time
type User struct {
	ID       uuid.UUID `json:"id"`
	UserName string    `json:"username"`
	Score    int       `json:"score"`
}

// Avoid exposing the uuid of the users in the response
// This will prevent users from impersonating other users
type UserResponse struct {
	UserName string `json:"username"`
	Score    int    `json:"score"`
}
