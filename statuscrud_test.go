package ssucore

import (
	"errors"
	"testing"
)

type statusCrudMockImplementer string

func (s statusCrudMockImplementer) GetStatuses() []Status {
	return []Status{
		Status{"test", "emojiA", "Status One"},
		Status{"test1", "emojiB", "Status Two"},
		Status{"test2", "emojiC", "Status Three"},
		Status{"test3", "emojiD", "Status Four"},
	}
}

func (s statusCrudMockImplementer) GetStatusByKey(key string) (Status, error) {
	if key == "test" {
		return Status{"test", "emojiA", "Status One"}, nil
	} else if key == "statusalreadyexisting" {
		return Status{"statusalreadyexisting", "emojiExisting", "An existing status"}, nil
	}
	return Status{}, errors.New("Could not find status with that key")
}

func (s statusCrudMockImplementer) AddNewStatus(status Status) error {
	if status.StatusName == "mynewstatus" {
		return nil
	}
	return nil
}

func TestStatusCRUDWillReturnStatuses(t *testing.T) {
	var sci StatusCRUDInterface
	sci = statusCrudMockImplementer("testing")
	statusList := GetStatuses(sci)
	if len(statusList) != 4 {
		t.Fatalf("Length of returned array was not %d. Expected 4", len(statusList))
	}
}

func TestStatusCRUDRunsSelectByKey(t *testing.T) {
	var sci StatusCRUDInterface
	sci = statusCrudMockImplementer("testing")
	status, err := GetStatusByKey("test", sci)
	if err != nil {
		t.Fatalf("Received an error when nil was expected")
	}
	if status.StatusName != "test" {
		t.Fatalf("Returned status did not match expected status name. Got %s, Expected %s", status.StatusName, "test")
	}
	status, err = GetStatusByKey("invalidkey", sci)
	if status.StatusName != "" {
		t.Fatalf("Received a status while expecting nil")
	}
	if err == nil {
		t.Fatalf("Did not receive expected error")
	}
	if err.Error() != "Could not find status with that key" {
		t.Fatalf("Received %s as error message. Expected %s", err.Error(), "Could not find status with that key")
	}
}

func TestStatusCRUDRunsAddStatusCorrectly(t *testing.T) {
	var sci StatusCRUDInterface
	sci = statusCrudMockImplementer("testing")
	err := AddNewStatus(Status{"mynewstatus", "emojiyey", "a hope new status"}, sci)
	if err != nil {
		t.Fatalf("Received %s as an error while expecting nil", err.Error())
	}

	err = AddNewStatus(Status{"statusalreadyexisting", "emojiyey", "a hope new status"}, sci)
	expectedErrorMessage := "Status already exists"
	if err == nil {
		t.Fatalf("Did not receive an error while expecting one")
	}
	if err.Error() != expectedErrorMessage {
		t.Fatalf("Received %s as error while expecting %s", err.Error(), expectedErrorMessage)
	}

	err = AddNewStatus(Status{"", "anemoji", ""}, sci)
	expectedErrorMessage = "Invalid Status (status without name) was supplied"
	if err == nil {
		t.Fatalf("Did not receive an error while expecting one")
	}
	if err.Error() != expectedErrorMessage {
		t.Fatalf("Received %s as error while expecting %s", err.Error(), expectedErrorMessage)
	}

	err = AddNewStatus(Status{"mynewstatus", "", ""}, sci)
	expectedErrorMessage = "Invalid Status (status with both text and emoji empty) was provided"
	if err == nil {
		t.Fatalf("Did not receive an error while expecting one")
	}
	if err.Error() != expectedErrorMessage {
		t.Fatalf("Received %s as error while expecting %s", err.Error(), expectedErrorMessage)
	}
}
