package types

type UsageUnit struct {
	Title string `json:"title" binding:"required"`
	Id    UUID   `json:"id" swaggerignore:"true"`
}

const (
	// time based
	UsageUnitHours = "hour"
	UsageUnitDays  = "day"

	// distance based
	UsageUnitMiles = "mile"
)
