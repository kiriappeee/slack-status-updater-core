package ssucore

import (
	"fmt"
	"strings"
	"errors"
	"gopkg.in/yaml.v2"
)

type Status struct {
	StatusName string `yaml:"statusName"`
	Emoji      string `yaml:"emoji,omitempty"`
	StatusText string `yaml:"statusText,omitempty"`
}

func ConvertTextToStructArray(textToConvert string) ([]Status, error) {

	var statusesToReturn []Status
	err := yaml.UnmarshalStrict([]byte(textToConvert), &statusesToReturn)
	if err != nil {
		if strings.Contains(err.Error(), "could not find expected") {
			fmt.Printf("Yaml was poorly formatted. Please check your config.\nError thrown was: %s\n", err)
		} else if strings.Contains(err.Error(), "unmarshal errors") {
			fmt.Printf("One or more fields in the yaml are incorrectly specified.\nError thrown was: %s\n", err)
		}
		return nil, err
	}

	checkedKeys := map[string]Status{}
	for _, status := range statusesToReturn {
		if status.StatusName == "" {
			return nil, errors.New(fmt.Sprintf("Status name (statusName) cannot be empty. Status in error is: %s", status.StatusName))
		}
		if status.Emoji == "" && status.StatusText == "" {
			return nil, errors.New(fmt.Sprintf("Both emoji and status text cannot be empty. Status in error is: %s", status.StatusName))
		}
		if checkedKeys[status.StatusName].StatusName != "" {
			return nil, errors.New(fmt.Sprintf("A duplicate key exists. Duplicate key found was %s", status.StatusName))
		}
		checkedKeys[status.StatusName] = status
	}
	return statusesToReturn, nil
}