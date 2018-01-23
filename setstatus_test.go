package ssucore

import "testing"

func TestSlackStatusCanBeInitializedCorrectly(t *testing.T){
	statusToTest := Status{statusName: "lunch", emoji: "chompy", status: "eating lunch"}
	if statusToTest.statusName != "lunch" {
		t.Errorf("Status name was %s. Expected lunch", statusToTest.statusName)
	}
}