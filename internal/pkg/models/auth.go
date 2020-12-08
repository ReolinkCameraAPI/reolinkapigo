package models

type LoginToken struct {
	Name      string `json:"name"`
	LeaseTime int    `json:"leaseTime"`
}

type Logout struct {
	RspCode int `json:"rspCode"`
}
