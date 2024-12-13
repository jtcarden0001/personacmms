package types

// a time unit is a descrete definition of a unit of time for the purposes of scheduling tasks
type TimeUnit struct {
	Title string `json:"title" binding:"required"`
	Id    UUID   `json:"id" swaggerignore:"true"`
}

const (
	TimeUnitDays   = "day"
	TimeUnitWeeks  = "week"
	TimeUnitMonths = "month"
	TimeUnitYears  = "year"
)
