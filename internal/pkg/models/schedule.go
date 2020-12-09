package models

type Schedule struct {
	Enable bool   `json:"enable"`
	Table  string `json:"table"`
}