package models

type OsdChannelData struct {
	Enable bool   `json:"enable"`
	Name   string `json:"name"`
	Pos    string `json:"pos"`
}

type OsdTimeData struct {
	Enable bool   `json:"enable"`
	Pos    string `json:"pos"`
}

type OsdData struct {
	BgColor    int            `json:"bgcolor"`
	Channel    int            `json:"channel"`
	OsdChannel OsdChannelData `json:"osdChannel"`
	OsdTime    OsdTimeData    `json:"osdTime"`
}
