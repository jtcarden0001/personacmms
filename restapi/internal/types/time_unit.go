package types

import "github.com/google/uuid"

// a time unit is a descrete definition of a unit of time for the purposes of scheduling tasks
type TimeUnit struct {
	Title string    `json:"title" binding:"required"`
	Id    uuid.UUID `json:"id" swaggerignore:"true"`
}

const (
	TimeUnitDays   = "day"
	TimeUnitWeeks  = "week"
	TimeUnitMonths = "month"
	TimeUnitYears  = "year"
)
