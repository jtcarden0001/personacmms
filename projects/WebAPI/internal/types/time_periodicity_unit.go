package types

type TimePeriodicityUnit struct {
	Id    int    `json:"id"`
	Title string `json:"title" binding:"required"`
}
