package prompt

import (
	"errors"
	"github.com/AlecAivazis/survey/v2"
)

func PromptString(prompt string) (string, error) {
	response := ""
	p := &survey.Input{
		Message: prompt,
	}
	survey.AskOne(p, &response)
	if response == "" {
		return "", errors.New("Empty response")
	}
	return response, nil
}
