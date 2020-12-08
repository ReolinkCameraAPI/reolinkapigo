package enum

type OsdPosition uint

const (
	UPPER_LEFT OsdPosition = iota
	TOP_CENTER
	UPPER_RIGHT
	LOWER_LEFT
	BOTTOM_CENTER
	LOWER_RIGHT
)

func (op OsdPosition) Name() string {
	return []string{"Upper Left", "Top Center", "Upper Right", "Lower Left", "Bottom Center", "Lower Right"}[op]
}
