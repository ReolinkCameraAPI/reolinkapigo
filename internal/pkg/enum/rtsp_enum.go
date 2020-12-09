package enum

type RtspProfile uint

const (
	BASE RtspProfile = iota
	MAIN
)

func (rp RtspProfile) Value() string {
	return []string{"base", "main"}[rp]
}
