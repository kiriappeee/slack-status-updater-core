package ssucore

import (
	"fmt"
	"strings"
	"gopkg.in/yaml.v2"
)

type Status struct {
	StatusName string `yaml:"statusName"`
	Emoji      string `yaml:"emoji"`
	StatusText string `yaml:"statusText"`
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

	return statusesToReturn, nil
}