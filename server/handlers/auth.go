package handlers

import (
	"github.com/BryanMwangi/pine"
	"github.com/BryanMwangi/qa/server/db/repo"
	"github.com/BryanMwangi/qa/server/utils"
	"github.com/google/uuid"
)

// this is an authentication middleware to authenticate the user
//
// the authentication can be done in many forms, including a cookies but for now
// let us use a header parameter
func authenticateRequest() pine.Handler {
	return func(c *pine.Ctx) error {
		if c.Header("Username") == "" {
			return c.SendStatus(401)
		}
		if c.Header("Authorization") == "" {
			return c.SendStatus(401)
		}
		// parse the uuid
		uuid, err := uuid.Parse(c.Header("Authorization"))
		if err != nil {
			return c.SendStatus(401)
		}
		pass, userName, err := utils.ValidateName(c.Header("Username"))
		if err != nil {
			return c.SendStatus(401)
		}
		if !pass {
			return c.SendStatus(401)
		}
		user := repo.GetUser(userName)
		if user.ID != uuid {
			return c.SendStatus(401)
		}
		c.Locals("userName", user.UserName)
		return c.Next()
	}
}
