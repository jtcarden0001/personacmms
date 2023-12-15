package types

type WorkOrderStatus struct {
	Id    int    `json:"id"`
	Title string `json:"title" binding:"required"`
}
