package models

type User struct {
	Level             string `json:"level"`
	Username          string `json:"userName"`
	CanBeDisconnected bool   `json:"canbeDisconn,omitempty"`
	IP                string `json:"ip,omitempty"`
	SessionId         int    `json:"sessionId,omitempty"`
}
