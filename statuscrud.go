package ssucore

type StatusCRUDInterface interface {
	List() []Status
}

func GetStatuses(sci StatusCRUDInterface) []Status {
	return sci.List()
}
