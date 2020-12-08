package models

type HddInfo struct {
	Capacity int `json:"capacity"`
	Format   int `json:"format"`
	ID       int `json:"id"`
	Mount    int `json:"mount"`
	Size     int `json:"size"`
}
