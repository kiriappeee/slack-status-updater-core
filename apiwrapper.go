package ssucore


type UpdateStatus func(*Status, string) (string, error)

func UpdateStatusViaSDK(s *Status, apiToken string) (string, error) {
	return "something", nil
}