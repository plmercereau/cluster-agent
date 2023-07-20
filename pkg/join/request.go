package join

import (
	"log"
	"net/rpc"

	"github.com/plmercereau/cluster-agent/pkg/mdns_service"
)

type JoinArgs struct{}

func JoinCluster() {
	// TODO ask the cluster node to join the cluster through the http server
	// TODO copy the necessary files e.g. /etc/corosync/corosync.conf
	// TODO run the pacemaker command to join the cluster once the cluster is configured

	address := mdns_service.ReachCluster()
	client, err := rpc.DialHTTP("tcp", address)
	if err != nil {
		log.Fatalln("Cannot dial HTTP", err)
	}

	var reply string
	args := JoinArgs{}

	err = client.Call("Kiosk.Join", args, &reply)
	if err != nil {
		log.Fatalln("Error calling service Kiosk.Join", err)
	}
	log.Println(reply)

}
