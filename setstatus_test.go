package ssucore

import (
	"testing"
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
	var statusesToTest []Status = ConvertTextToStructArray(textToReadIn)
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
