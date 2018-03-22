package ssucore

import (
	"fmt"
	"github.com/nlopes/slack"
)

type UpdateStatus func(*Status, string) (string, error)

func UpdateStatusViaSDK(s *Status, apiToken string) (string, error) {
	api := slack.New(apiToken)
	err := api.SetUserCustomStatus(s.StatusText, fmt.Sprintf(":%s:",s.Emoji))
	if err == nil{
	    return fmt.Sprintf("Status set for %s", s.StatusName), nil
	} else {
		return fmt.Sprintf("Status set failed for %s", s.StatusName), err
	}
}