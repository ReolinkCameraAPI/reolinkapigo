package network

type Scheme uint

const (
	HTTP Scheme = iota
	HTTPS
	SOCKS5
)

func (p Scheme) String() string {
	return []string{"http", "https", "socks5"}[p]
}

type Protocol uint

const (
	PROTOCOL_UDP Protocol = iota
	PROTOCOL_TCP
)

func (p Protocol) String() string {
	return []string{"udp", "tcp"}[p]
}
