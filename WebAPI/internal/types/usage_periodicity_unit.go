package types

type UsagePeriodicityUnit struct {
	Id    int    `json:"id"`
	Title string `json:"title" binding:"required"`
}
