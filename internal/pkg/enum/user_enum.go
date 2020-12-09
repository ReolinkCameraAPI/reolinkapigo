package enum

type UserLevel uint

const (
	GUEST UserLevel = iota
	ADMIN
)

func (ul UserLevel) Value() string {
	return []string{"guest", "admin"}[ul]
}
