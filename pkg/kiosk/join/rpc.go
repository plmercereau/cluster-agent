package join

import (
	"log"

	"github.com/plmercereau/cluster-agent/pkg/agent/gate/close"
	"github.com/plmercereau/cluster-agent/pkg/pacemaker"
	"github.com/plmercereau/cluster-agent/pkg/timer"
)

// Open the cluster node so a new node can join
func AddNode(args JoinArgs) {
	timer.CancelTimer()
	log.Println("Handling a request to join the cluster...")

	err := pacemaker.ActivateMaintenanceMode()
	if err != nil {
		log.Fatalln("Error while activating maintenance mode.", err)
	}

	pacemaker.AuthenticateNode(args.Host, args.IP, args.Username, args.Password)
	pacemaker.AddNode(args.Host)
	pacemaker.DeactivateMaintenanceMode()
	log.Println("Node X added to the cluster.")
	close.CloseGate()
}
