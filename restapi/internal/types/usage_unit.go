package types

type UsageUnit struct {
	Title string `json:"title" binding:"required"`
	Id    UUID   `json:"id" swaggerignore:"true"`
	Type  string `json:"type" binding:"required"`
}

const (
	// types
	UsageUnitTypeTime     = "time"
	UsageUnitTypeDistance = "distance"

	// units
	// time based
	UsageUnitHours = "hour"
	UsageUnitDays  = "day"

	// distance based
	UsageUnitMiles = "mile"
)
