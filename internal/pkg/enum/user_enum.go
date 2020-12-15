package enum

type UserLevel uint

const (
	USER_LEVEL_GUEST UserLevel = iota
	USER_LEVEL_ADMIN
)

func (ul UserLevel) Value() string {
	return []string{"guest", "admin"}[ul]
}
