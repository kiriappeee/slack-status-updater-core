package ssucore

import (
	"testing"
)

type StatusCrudMockImplementer string

func (s StatusCrudMockImplementer) List() []Status {
	return []Status{
		Status{"test", "me", "now"},
		Status{"test", "me", "now"},
		Status{"test", "me", "now"},
		Status{"test", "me", "now"},
	}
}
func TestStatusCRUDWillReturnStatuses(t *testing.T) {
	var sci StatusCRUDInterface
	sci = StatusCrudMockImplementer("testing")
	statusList := GetStatuses(sci)
	if len(statusList) != 4 {
		t.Fatalf("Length of returned array was not %d. Expected 4", len(statusList))
	}
}
