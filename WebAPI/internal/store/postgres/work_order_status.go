package postgres

import tp "github.com/jtcarden0001/personacmms/webapi/internal/types"

type WorkOrderStatus interface {
	CreateWorkOrderStatus(string) (int, error)
	DeleteWorkOrderStatus(int) error
	GetAllWorkOrderStatus() ([]tp.WorkOrderStatus, error)
	GetWorkOrderStatus(int) (tp.WorkOrderStatus, error)
	UpdateWorkOrderStatus(int, string) error
}

type WorkOrderStatusTest interface {
	ResetSequenceWorkOrderStatus(int) error
}
