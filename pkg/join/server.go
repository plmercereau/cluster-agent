package join

import (
	"log"

	"github.com/plmercereau/cluster-agent/pkg/agent/gate/close"
	"github.com/plmercereau/cluster-agent/pkg/timer"
)

// Open the cluster node so a new node can join
func AddNode() {
	timer.CancelTimer()
	log.Println("Handling a request to join the cluster...")

	log.Println("Node X added to the cluster.")
	close.CloseGate()

}
