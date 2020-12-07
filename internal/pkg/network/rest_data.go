package network

type GeneralData struct {
	Code    int         `json:"code"`
	Value   interface{} `json:"value;omitempty"`
	Initial interface{} `json:"initial;omitempty"`
	Range   interface{} `json:"range;omitempty"`
}
