package rest

import "encoding/json"

type GeneralData struct {
	Cmd     string                     `json:"cmd"`
	Code    int                        `json:"code"`
	Value   map[string]json.RawMessage `json:"value,omitempty"`
	Initial map[string]json.RawMessage `json:"initial,omitempty"`
	Range   map[string]json.RawMessage `json:"range,omitempty"`
}
