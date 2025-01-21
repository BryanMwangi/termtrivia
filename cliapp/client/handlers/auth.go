package handlers

import (
	"encoding/json"
	"time"

	"github.com/BryanMwangi/qa/cliapp/client"
	"github.com/BryanMwangi/qa/cliapp/client/models"
	"github.com/BryanMwangi/qa/cliapp/utils"
)

func FetchUser(userName string) error {
	userUri := client.GenerateUri("/user/" + userName)
	client.Client.Request().SetRequestURI(userUri)

	if err := client.Client.SendRequest(); err != nil {
		return err
	}

	utils.ShowLoader("Authenticating user...")
	_, body, err := client.Client.ReadResponse()
	if err != nil {
		return err
	}
	var user models.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		return err
	}
	headers := map[string]string{
		"Authorization": user.ID,
		"Username":      user.UserName,
	}

	// Once authenticated the client is globally authenticated since the headers
	// are set on the client
	//
	// This is the only part of Pine that allows for the state to remain as is
	// even after sending a request promoting reusability
	client.Client.Request().SetHeaders(headers)

	// small delay to improve UX
	utils.StopLoader(3 * time.Second)
	return nil
}
