package ssucore

type UpdateStatus func(*Status) (string, error)

func UpdateStatusViaSDK(s *Status) (string, error) {
	return "something", nil
}