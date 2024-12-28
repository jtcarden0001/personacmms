package types

import (
	"strings"

	"github.com/google/uuid"
)

// a time trigger is an event that is triggered after a specific time has elapsed since the last time a
// work order was completed for a task.
type TimeTrigger struct {
	Id       uuid.UUID `json:"id" swaggerignore:"true"`
	Quantity int       `json:"quantity" binding:"required"`
	TimeUnit string    `json:"time_unit" binding:"required"`
	TaskId   uuid.UUID `json:"asset_task_id" swaggerignore:"true"`
}

const (
	TimeTriggerUnitDays   = "day"
	TimeTriggerUnitWeeks  = "week"
	TimeTriggerUnitMonths = "month"
	TimeTriggerUnitYears  = "year"
)

var ValidTimeTriggerUnits = map[string]bool{
	TimeTriggerUnitDays:   true,
	TimeTriggerUnitWeeks:  true,
	TimeTriggerUnitMonths: true,
	TimeTriggerUnitYears:  true,
}

func PrintValidTimeTriggerUnits() string {
	units := []string{}
	for unit := range ValidTimeTriggerUnits {
		units = append(units, unit)
	}
	return strings.Join(units, ", ")
}
