package models

// TODO: update
type PtzOperation struct {}

type PtzPreset struct{
	Channel int    `json:"channel"`
	Enable  int    `json:"enable"`
	Index   int    `json:"id"`
	Name    string `json:"name"`
}
