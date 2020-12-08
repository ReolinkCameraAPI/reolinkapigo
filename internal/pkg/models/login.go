package models

type LoginToken struct {
	Name      string `json:"name"`
	LeaseTime int    `json:"leaseTime"`
}
