package types

type TimeUnit struct {
	Id    int    `json:"id"`
	Title string `json:"title" binding:"required"`
}
