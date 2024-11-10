package types

type AssetCategory struct {
	Id    int    `json:"id"`
	Title string `json:"title" binding:"required"`
}
