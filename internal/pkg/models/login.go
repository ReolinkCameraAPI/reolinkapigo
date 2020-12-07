package models

type LoginTokenData struct {
	Name string `json:"name"`
}

type LoginData struct {
	Token LoginTokenData `json:"Token"`
}
