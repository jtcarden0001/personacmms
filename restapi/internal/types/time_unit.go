package types

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
