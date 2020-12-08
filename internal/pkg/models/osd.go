package models

type OsdChannel struct {
	Enable bool   `json:"enable"`
	Name   string `json:"name"`
	Pos    string `json:"pos"`
}

type OsdTime struct {
	Enable bool   `json:"enable"`
	Pos    string `json:"pos"`
}

type Osd struct {
	BgColor    bool       `json:"bgcolor"`
	Channel    int        `json:"channel"`
	OsdChannel OsdChannel `json:"osdChannel"`
	OsdTime    OsdTime    `json:"osdTime"`
}
