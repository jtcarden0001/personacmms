package types

import tm "time"

type WorkOrder struct {
	Id                 int      `json:"id"`
	PreventativeTaskId int      `json:"preventativeTaskId" binding:"required"`
	StatusId           int      `json:"statusId" binding:"required"`
	CreatedDate        tm.Time  `json:"createdDate" binding:"required"`
	CompletedDate      *tm.Time `json:"completedDate"`
}
