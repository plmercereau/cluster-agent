package kiosk

import "github.com/plmercereau/cluster-agent/pkg/kiosk/join"

type Kiosk struct {
}

func (g Kiosk) Join(_ join.JoinArgs, reply *string) error {
	join.AddNode()
	*reply = "Joined cluster X."
	return nil
}
