package enum

// === Image Settings enums ===
// TODO: still need to finalise this

// ===
type AntiFlicker uint

const (
	OUTDOOR AntiFlicker = iota
)

func (i AntiFlicker) Value() string {
	return []string{"Outdoor"}[i]
}
// ===

type Exposure uint

const (
	AUTO Exposure = iota
)

func (i Exposure) Value() string {
	return []string{"Auto"}[i]
}
// ===