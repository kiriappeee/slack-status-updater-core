package ssucore

import (
	"log"

	"gopkg.in/yaml.v2"
)

type Status struct {
	StatusName string `yaml:"statusName"`
	Emoji      string `yaml:"emoji"`
	StatusText string `yaml:"statusText"`
}

func ConvertTextToStructArray(textToConvert string) []Status {

	var statusesToReturn []Status
	err := yaml.Unmarshal([]byte(textToConvert), &statusesToReturn)
	if err != nil {
		log.Fatalf("couldn't convert it %s", err)
	}

	return statusesToReturn
}
