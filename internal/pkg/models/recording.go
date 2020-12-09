package models

type RecordingMainStream struct {
	BitRate   int    `json:"bitRate"`
	FrameRate int    `json:"frameRate"`
	Profile   string `json:"profile"`
	Size      string `json:"size"`
}

type RecordingSubStream struct {
	BitRate   int    `json:"bitRate"`
	FrameRate int    `json:"frameRate"`
	Profile   string `json:"profile"`
	Size      string `json:"size"`
}

// TODO: update
type Encoding struct {
	Audio      bool                `json:"audio"`
	Channel    int                 `json:"channel"`
	MainStream RecordingMainStream `json:"mainStream"`
	SubStream  RecordingSubStream  `json:"subStream"`
}

type Recording struct {
	Channel    int      `json:"channel"`
	Overwrite  bool     `json:"overwrite"`
	PostRecord string   `json:"postRec"`
	PreRecord  bool     `json:"preRec"`
	Schedule   Schedule `json:"schedule"`
}
