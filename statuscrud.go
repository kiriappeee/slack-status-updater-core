package ssucore

import (
	"errors"
)

type StatusCRUDInterface interface {
	GetStatuses() []Status
	GetStatusByKey(k string) (Status, error)
	AddNewStatus(status Status) error
	DeleteStatusByName(nameOfStatusToDelete string) error
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

func DeleteStatusByName(nameOfStatusToDelete string, sci StatusCRUDInterface) error {
	_, err := sci.GetStatusByKey(nameOfStatusToDelete)
	if err != nil {
		return err
	}
	return sci.DeleteStatusByName(nameOfStatusToDelete)
}

func EditStatus(nameOfStatusBeingEdited string, editedStatus Status, sci StatusCRUDInterface) error {
	if editedStatus.StatusName == "" {
		return errors.New("Invalid Status (status without name) was supplied")
	}

	if editedStatus.StatusText == "" && editedStatus.Emoji == "" {
		return errors.New("Invalid Status (status with both text and emoji empty) was provided")
	}

	if nameOfStatusBeingEdited != editedStatus.StatusName {
		_, err := sci.GetStatusByKey(editedStatus.StatusName)
		if err == nil {
			return errors.New("Status with that name already exists")
		}
	}

	err := sci.DeleteStatusByName(nameOfStatusBeingEdited)
	if err != nil {
		return err
	}

	err = sci.AddNewStatus(editedStatus)
	if err != nil {
		return err
	}
	return nil
}
