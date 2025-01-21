package cmd

import (
	"fmt"

	"github.com/BryanMwangi/qa/cliapp/client/handlers"
	"github.com/BryanMwangi/qa/cliapp/utils"
	"github.com/manifoldco/promptui"
)

func authenticateUser() error {
	validate := func(input string) error {
		pass, _, err := utils.ValidateName(input)
		if err != nil {
			return err
		}
		if !pass {
			return fmt.Errorf("name contains illegal characters")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Username",
		Validate: validate,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return err
	}
	_, name, err := utils.ValidateName(result)
	if err != nil {
		fmt.Printf("Invalid username %v\n", err)
		return err
	}

	err = handlers.FetchUser(name)
	if err != nil {
		fmt.Printf("Failed to fetch user %v\n", err)
		return err
	}
	fmt.Println("Welcome ", name)
	return nil
}
