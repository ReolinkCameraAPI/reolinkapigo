package models

type GetHddInfoData struct {
	Capacity int `json:"capacity"`
	Format   int `json:"format"`
	ID       int `json:"id"`
	Mount    int `json:"mount"`
	Size     int `json:"size"`
}
