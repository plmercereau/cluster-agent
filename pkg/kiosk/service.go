package kiosk

import (
	"github.com/plmercereau/cluster-agent/pkg/kiosk/join"
)

type Kiosk struct {
}

func (g Kiosk) Join(args join.JoinArgs, reply *string) error {
	join.AddNode(args)
	*reply = "OK"
	return nil
}
