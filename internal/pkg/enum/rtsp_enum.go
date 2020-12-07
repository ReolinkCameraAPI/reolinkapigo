package enum

type RtspProfile uint

const (
	BASE RtspProfile = iota
	MAIN
)

func (rp RtspProfile) String() string {
	return []string{"base", "main"}[rp]
}
