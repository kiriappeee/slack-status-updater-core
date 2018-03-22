package ssucore

import (
	"testing"
	"os"
)

func TestSlackStatusUpdateReturnsCorrectMessage(t *testing.T){
	s := Status{"lunch", "chompy", "having lunch"}
	_, err := UpdateStatusViaSDK(&s, os.Getenv("SLACKTOKEN"))
	if err != nil{
		t.Fatalf("Error was not nil. Received %s", err.Error())
	}
}