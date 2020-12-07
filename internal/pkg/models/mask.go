package models

type MaskAreaBlock struct {
	Height int `json:"height"`
	Width  int `json:"width"`
	X      int `json:"x"`
	Y      int `json:"y"`
}

type MaskAreaScreen struct {
	Height int `json:"height"`
	Width  int `json:"width"`
}

type MaskArea struct {
	Block  MaskAreaBlock  `json:"block"`
	Screen MaskAreaScreen `json:"screen"`
}

type MaskData struct {
	Area    []MaskArea `json:"area"`
	Channel int        `json:"channel"`
	Enable  bool       `json:"enable"`
}