package types

type Task struct {
	Id                       int    `json:"id"`
	Title                    string `json:"title" binding:"required"`
	Instructions             string `json:"instructions" binding:"required"`
	TimePeriodicityQuantity  *int   `json:"timePeriodicityQuantity"`
	TimePeriodicityUnitId    *int   `json:"timePeriodicityUnitId"`
	UsagePeriodicityQuantity *int   `json:"usagePeriodicityQuantity"`
	UsagePeriodicityUnitId   *int   `json:"usagePeriodicityUnitId"`
	EquipmentId              int    `json:"equipmentId"`
}
