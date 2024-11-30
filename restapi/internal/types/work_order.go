package types

import tm "time"

type WorkOrder struct {
	Id              UUID     `json:"id" swaggerignore:"true"`
	CreatedDate     tm.Time  `json:"createdDate" swaggerignore:"true"`
	CompletedDate   *tm.Time `json:"completedDate"`
	Notes           *string  `json:"notes"`
	CumulativeMiles *int     `json:"cumulativeMiles"`
	CumulativeHours *int     `json:"cumulativeHours"`
	TaskId          UUID     `json:"TaskId" swaggerignore:"true"`
	StatusTitle     string   `json:"statusId" binding:"required"`
}
