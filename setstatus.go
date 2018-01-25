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

func ConvertTextToStructArray(textToConvert string) ([]Status, error) {

	var statusesToReturn []Status
	err := yaml.Unmarshal([]byte(textToConvert), &statusesToReturn)
	if err != nil {
		log.Printf("ssucore.ConvertTextToStructArray: couldn't convert it %s", err)
		return nil, err
	}

	return statusesToReturn, nil
}