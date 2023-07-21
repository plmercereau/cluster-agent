package join

import "net"

type JoinArgs struct {
	IP       net.IP
	Host     string
	Username string
	Password string
}
