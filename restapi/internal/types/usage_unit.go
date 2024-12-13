package types

// A UsageUnit is a unit of measurement that is used to measure the usage of a particular asset
// for the purpose of task scheduling
type UsageUnit struct {
	Title string `json:"title" binding:"required"`
	Id    UUID   `json:"id" swaggerignore:"true"`
	Type  string `json:"type" binding:"required"`
}

// If you add something here then add it below in the validUsageUnitTypes map
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

var ValidUsageUnitTypes = map[string]bool{
	UsageUnitTypeTime:     true,
	UsageUnitTypeDistance: true,
	UsageUnitHours:        true,
	UsageUnitDays:         true,
	UsageUnitMiles:        true,
}
