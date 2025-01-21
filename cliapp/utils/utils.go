package utils

import (
	"fmt"
	"regexp"
	"time"

	"github.com/briandowns/spinner"
)

var (
	TermSpinner *spinner.Spinner
)

func ValidateName(name string) (bool, string, error) {
	if name == "" {
		return false, "", fmt.Errorf("name is empty")
	}
	// Define the regex patterns
	nameRegex := `^[A-Za-z\s'-À-ÖØ-öø-ÿ]+$`
	sanitizerRegex := `[\/\?<>\\:\*\|":]`

	// Compile the regex for illegal characters
	sanitizerRe, err := regexp.Compile(sanitizerRegex)
	if err != nil {
		return false, "", fmt.Errorf("failed to compile sanitizer regex: %v", err)
	}

	// Remove illegal characters
	sanitizedString := sanitizerRe.ReplaceAllString(name, "")

	// Compile the regex for name validation
	re, err := regexp.Compile(nameRegex)
	if err != nil {
		return false, "", fmt.Errorf("failed to compile name regex: %v", err)
	}

	if re.MatchString(sanitizedString) {
		// Check if the sanitized string is empty
		if sanitizedString == "" {
			return false, "", fmt.Errorf("name contains only illegal characters")
		}
		return true, sanitizedString, nil
	} else {
		return false, "", fmt.Errorf("name contains only illegal characters")
	}
}

func ShowLoader(message string) {
	fmt.Println(message)
	TermSpinner = spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	TermSpinner.Start()
}

// This is used to stop the loader after a short delay
//
// Please note that this is not good practice because if the TermSpinner is not
// initialized, it will cause a panic
//
// I just did this for the sake of simplicity and to reduce overall memory usage
func StopLoader(sleep time.Duration) {
	time.Sleep(sleep)
	TermSpinner.Stop()
}
