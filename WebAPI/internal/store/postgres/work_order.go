package postgres

import tp "github.com/jtcarden0001/personacmms/webapi/internal/types"

type WorkOrder interface {
	CreateWorkOrder(int, int, string, *string) (int, error)
	DeleteWorkOrder(int) error
	GetAllWorkOrder() ([]tp.WorkOrder, error)
	GetAllWorkOrderByEquipmentId(int) ([]tp.WorkOrder, error)
	GetWorkOrder(int) (tp.WorkOrder, error)
	UpdateWorkOrder(int, int, int, string, *string) error
}

type WorkOrderTest interface {
	ResetSequenceWorkOrder(int) error
}
