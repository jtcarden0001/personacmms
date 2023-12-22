package types

import tm "time"

type WorkOrder struct {
	Id            int      `json:"id"`
	TaskId        int      `json:"taskId" binding:"required"`
	StatusId      int      `json:"statusId" binding:"required"`
	CreatedDate   tm.Time  `json:"createdDate"`
	CompletedDate *tm.Time `json:"completedDate"`
}
