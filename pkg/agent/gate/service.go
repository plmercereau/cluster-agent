package gate

import (
	"log"

	"github.com/plmercereau/cluster-agent/pkg/agent/gate/close"
	"github.com/plmercereau/cluster-agent/pkg/agent/gate/open"
	"github.com/plmercereau/cluster-agent/pkg/http_server"
)

type Gate struct {
}

func (g Gate) Close(_ close.CloseArgs, reply *string) error {
	if !http_server.IsLive() {
		log.Println("Gate is already closed.")
	} else {
		close.CloseGate()
	}
	*reply = "OK"
	return nil

}

// Open the cluster node so a new node can join
func (g Gate) Open(_ open.OpenArgs, reply *string) error {
	open.OpenGate()

	*reply = "The gate to join the cluster is open."
	return nil
}
