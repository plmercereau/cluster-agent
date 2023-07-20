package kiosk

import (
	"fmt"

	"github.com/plmercereau/cluster-agent/pkg/kiosk/join"
)

type Kiosk struct {
}

func (g Kiosk) Join(args join.JoinArgs, reply *string) error {
	join.AddNode(args.Host, args.Username, args.Password)
	*reply = fmt.Sprintf("Node %s added to the cluster.", args.Host)
	return nil
}
