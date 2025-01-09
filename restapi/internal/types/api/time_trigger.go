package api

import (
	"strings"

	"github.com/google/uuid"
)

// a time trigger is an event that is triggered after a specific time has elapsed since the last time a
// work order was completed for a task.
type TimeTriggerRequest struct {
	Quantity int    `json:"quantity" binding:"required"`
	TimeUnit string `json:"timeUnit" binding:"required"`
}

type TimeTriggerResponse struct {
	Id       uuid.UUID `json:"id"`
	Quantity int       `json:"quantity"`
	TimeUnit string    `json:"timeUnit"`
	TaskId   uuid.UUID `json:"-"`
	// task api reference "/api/v1/assets/{assetId}/tasks/{taskId}"
	TaskReference string `json:"taskReference"`
}

// TODO: define these in one place, currently duplicated in store and api t
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
