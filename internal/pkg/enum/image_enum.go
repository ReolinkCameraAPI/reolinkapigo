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
	EXPOSURE_AUTO Exposure = iota
)

func (i Exposure) Value() string {
	return []string{"Auto"}[i]
}

// ===

type Backlight uint

const (
	DYNAMIC_RANGE_CONTROL Backlight = iota
)

func (b Backlight) Value() string {
	return []string{"DynamicRangeControl"}[b]
}

// ==

type WhiteBalance uint

const (
	WHITE_BALANCE_AUTO WhiteBalance = iota
)

func (wb WhiteBalance) Value() string {
	return []string{"Auto"}[wb]
}

type DayNight uint

const (
	DAY_NIGHT_AUTO DayNight = iota
)

func (dn DayNight) Value() string {
	return []string{"Auto"}[dn]
}

type Rotation uint

const (
	ROTATION_0 Rotation = iota
	ROTATION_90
	ROTATION_180
	ROTATION_270
)

func (r Rotation) Value() int {
	return []int{1}[r]
}
