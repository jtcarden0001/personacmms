package types

type EquipmentCategory struct {
	Id    int    `json:"id"`
	Title string `json:"title" binding:"required"`
}
