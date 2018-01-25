package ssucore

import (
	"testing"
	"strings"
)

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
