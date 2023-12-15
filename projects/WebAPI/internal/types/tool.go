package types

type Tool struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Size  string `json:"size"`
	// TODO: might be nice to add an image of the tool
}
