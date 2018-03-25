package ssucore

import (
	"testing"
	"strings"
	"errors"
)

func updateStatusViaAPIMock(s *Status, apiToken string) (string, error){
	if s.StatusName == "lunch" && s.Emoji == "chompy" && s.StatusText == "Having lunch"{
		return "Status was successfully changed", nil
	} else {
		return "", errors.New("Could not update the status message")
	}
}

func TestSlackStatusCanBeInitializedCorrectly(t *testing.T) {
	var statusToTest Status = Status{StatusName: "lunch", Emoji: "chompy", StatusText: "eating lunch"}
	if statusToTest.StatusName != "lunch" {
		t.Fatalf("Status name was %s. Expected lunch", statusToTest.StatusName)
	}
}

func TestSlackStatusCanBeReadCorrectlyFromString(t *testing.T) {
	textToReadIn := `
- statusName: lunch
  emoji: chompy
  statusText: Having lunch

- statusName: resting
  emoji: bath
  statusText: Resting

- statusName: awesome
  emoji: awesome
  statusText: ''

- statusName: deep-work
  emoji: ''
  statusText: In Focus mode
`
	statusesToTest, err := ConvertTextToStructArray(textToReadIn)
	if err != nil {
		t.Fatalf("An error occured when converting text. Did not expect error. Error was %v", err)
	}
	if len(statusesToTest) != 4 {
		t.Fatalf("Length of returned array was %d. Expected 4", len(statusesToTest))
	} else {
		var statusToTest Status = statusesToTest[0]
		if statusToTest.StatusName != "lunch" {
			t.Fatalf("Error when checking status 0: Status name was %s. Expected lunch", statusToTest.StatusName)
		}
		if statusToTest.Emoji != "chompy" {
			t.Fatalf("Error when checking status 0: Emoji was %s. Expected chompy", statusToTest.Emoji)
		}
		if statusToTest.StatusText != "Having lunch" {
			t.Fatalf("Error when checking status 0: Status text was %s. Expected Having lunch", statusToTest.StatusText)
		}
	}

}

func TestAnErrorIsThrownWhenAPoorlyFormattedYamlIsProvided(t *testing.T) {
	textToReadIn :=`
- statusName: lunch
  emoji: chompy
  statusText: Having lunch

- statusName: resting
  emoji 
  statusText: Resting

- statusName: awesome
   emoji: awesome
  statusText: ''
 -
`
	statusesToTest, err := ConvertTextToStructArray(textToReadIn)
	if err == nil {
		t.Fatalf("Error was nil. Expected an error")
	}

	if ! strings.Contains(err.Error(), "could not find expected") {
		t.Fatalf("Error message was not the expected one. Received: %s", err.Error())
	}

	if statusesToTest != nil{
		t.Fatalf("Statuses to test was %v. Expected nil", statusesToTest)
	}

}

func TestAnErrorIsThrownWhenTheYamlDoesntMatchTheAppSpec(t *testing.T){
	textToReadIn :=`
- statusName: lunch
  emoji: chompy
  statusText: Having lunch

- statusname: resting
  emoji: sleep
  statusText: Resting
`

	statusesToTest, err := ConvertTextToStructArray(textToReadIn)
	if err == nil {
		t.Fatalf("Error was nil. Expected an error")
	} else {
		if ! strings.Contains(err.Error(), "unmarshal errors") {
			t.Fatalf("Expected unmarshal errors error message. Received: %s", err.Error())
		}
	}
	if statusesToTest != nil{
		t.Fatalf("Statuses to test was %v. Expected nil", statusesToTest)
	}
}

func TestAnErrorIsThrownIfCorrectValuesAreNotSuppliedForTheStruct(t *testing.T){
	textToReadIn := `
- statusName: lunch
  emoji: ''
  statusText: ''
`
	
	_, err := ConvertTextToStructArray(textToReadIn)
	if err == nil {
		t.Fatalf("Error was nil. Expected an error")
	} else {
		if ! strings.Contains(err.Error(), "Both emoji and status text cannot be empty") {
			t.Fatalf("Expected error regarding missing values. Received: %s", err.Error())
		}
	}

	textToReadIn = `
- statusName: ''
  emoji: ''
  statusText: ''
`
	_, err = ConvertTextToStructArray(textToReadIn)
	if err == nil {
		t.Fatalf("Error was nil. Expected an error")
	} else {
		if ! strings.Contains(err.Error(), "Status name (statusName) cannot be empty. ") {
			t.Fatalf("Expected error regarding missing status name. Received: %s", err.Error())
		}
	}

	textToReadIn = `
- statusName: lunch
  statusText: ''
`
	
	_, err = ConvertTextToStructArray(textToReadIn)
	if err == nil {
		t.Fatalf("Error was nil. Expected an error")
	} else {
		if ! strings.Contains(err.Error(), "Both emoji and status text cannot be empty") {
			t.Fatalf("Expected error regarding missing values. Received: %s", err.Error())
		}
	}

	textToReadIn = `
- statusName: lunch
  emoji: ''
`
	
	_, err = ConvertTextToStructArray(textToReadIn)
	if err == nil {
		t.Fatalf("Error was nil. Expected an error")
	} else {
		if ! strings.Contains(err.Error(), "Both emoji and status text cannot be empty") {
			t.Fatalf("Expected error regarding missing values. Received: %s", err.Error())
		}
	}

	textToReadIn = `
- statusName: lunch'
`
	
	_, err = ConvertTextToStructArray(textToReadIn)
	if err == nil {
		t.Fatalf("Error was nil. Expected an error")
	} else {
		if ! strings.Contains(err.Error(), "Both emoji and status text cannot be empty") {
			t.Fatalf("Expected error regarding missing values. Received: %s", err.Error())
		}
	}
}

func TestAnErrorIsThrownWhenADuplicateValueExists(t *testing.T){
	textToReadIn :=`
- statusName: lunch
  emoji: chompy
  statusText: Having lunch

- statusName: resting
  emoji: bath
  statusText: Resting

- statusName: lunch
  emoji: awesome
  statusText: Having lunch
`
	_, err := ConvertTextToStructArray(textToReadIn)
	if err == nil {
		t.Fatalf("Error was nil. Expected an error")
	} else {
		if ! strings.Contains(err.Error(), "duplicate key") {
			t.Fatalf("Expected duplicate key error message. Received: %s", err.Error())
		}
	}
}

func TestSetStatusMethodIsCalled(t *testing.T){
	s := Status{"lunch", "chompy", "Having lunch"}
	result, err := s.SetMyStatus(updateStatusViaAPIMock, "mytoken")
	if  err != nil {
		t.Fatalf("Error was not nil. Received: %s", err.Error())
	}
	if result != "Status was successfully changed" {
		t.Fatalf("Result message was not the expected value. Received: %s", result)
	}

	s = Status{"lunch", "chompy", "eating"}
	result, err = s.SetMyStatus(updateStatusViaAPIMock, "mytoken")
	if  err == nil {
		t.Fatalf("Error was nil. Expected an error")
	}

	if err.Error() != "Could not update the status message" {
		t.Fatalf("Received error message was not as expected. Received: %s", err.Error())
	}
}