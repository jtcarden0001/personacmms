package types

type WorkOrder struct {
	Id            int     `json:"id"`
	TaskId        int     `json:"taskId" binding:"required"`
	StatusId      int     `json:"statusId" binding:"required"`
	CreatedDate   string  `json:"createdDate"`
	CompletedDate *string `json:"completedDate"`
}
