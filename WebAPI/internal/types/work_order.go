package types

type WorkOrder struct {
	Id            int     `json:"id"`
	EquipmentId   int     `json:"equipmentId" binding:"required"`
	TaskId        int     `json:"taskId" binding:"required"`
	StatusId      int     `json:"statusId" binding:"required"`
	CreatedDate   string  `json:"createdDate"`
	CompletedDate *string `json:"completedDate"`
}
