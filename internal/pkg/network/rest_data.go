package network

type GeneralData struct {
	Code    int         `json:"code"`
	Value   interface{} `json:"value;omitempty"`
	Initial interface{} `json:"initial;omitempty"`
	Range   interface{} `json:"range;omitempty"`
}

type LoginTokenData struct {
	Name string `json:"name"`
}

type LoginData struct {
	Token LoginTokenData `json:"Token"`
}

type GetHddInfoData struct {
	Capacity int `json:"capacity"`
	Format   int `json:"format"`
	ID       int `json:"id"`
	Mount    int `json:"mount"`
	Size     int `json:"size"`
}

type FormatHddData struct {
	RspCode int `json:"rspCode"`
}

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
