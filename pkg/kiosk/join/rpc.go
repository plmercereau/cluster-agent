package join

import (
	"log"

	"github.com/plmercereau/cluster-agent/pkg/agent/gate/close"
	"github.com/plmercereau/cluster-agent/pkg/pacemaker"
	"github.com/plmercereau/cluster-agent/pkg/timer"
)

// Open the cluster node so a new node can join
func AddNode(host string, username string, password string) {
	timer.CancelTimer()
	log.Println("Handling a request to join the cluster...")
	// TODO put the cluster into maintenance mode
	err := pacemaker.ActivateMaintenanceMode()
	if err != nil {
		log.Fatalln("Error while activating maintenance mode.")
	}

	pacemaker.AuthenticateNode(host, username, password)
	pacemaker.AddNode(host)
	pacemaker.DeactivateMaintenanceMode()
	log.Println("Node X added to the cluster.")
	close.CloseGate()

}
