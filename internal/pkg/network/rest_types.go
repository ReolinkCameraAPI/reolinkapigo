package network

type Proxy uint

const (
	HTTP Proxy = iota
	HTTPS
	SOCKS5
)

func (p Proxy) String() string {
	return []string{"http", "https", "socks5"}[p]
}

type Protocol uint

const (
	UDP Protocol = iota
	TCP
)

func (p Protocol) String() string {
	return []string{"udp", "tcp"}[p]
}
