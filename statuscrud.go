package ssucore

import (
	"errors"
)

type StatusCRUDInterface interface {
	GetStatuses() []Status
	GetStatusByKey(k string) (Status, error)
	AddNewStatus(status Status) error
}

func GetStatuses(sci StatusCRUDInterface) []Status {
	return sci.GetStatuses()
}

func GetStatusByKey(key string, sci StatusCRUDInterface) (Status, error) {
	return sci.GetStatusByKey(key)
}

func AddNewStatus(status Status, sci StatusCRUDInterface) error {
	if status.StatusName == "" {
		return errors.New("Invalid Status (status without name) was supplied")
	}
	if status.StatusText == "" && status.Emoji == "" {
		return errors.New("Invalid Status (status with both text and emoji empty) was provided")
	}
	_, err := sci.GetStatusByKey(status.StatusName)
	if err == nil {
		return errors.New("Status already exists")
	}
	return sci.AddNewStatus(status)
}
