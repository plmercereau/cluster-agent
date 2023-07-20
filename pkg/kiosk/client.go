package kiosk

import (
	"log"
	"net/rpc"

	"github.com/plmercereau/cluster-agent/pkg/kiosk/join"
	"github.com/plmercereau/cluster-agent/pkg/mdns_service"
	"github.com/plmercereau/cluster-agent/pkg/pacemaker"
)

type KioskClient struct {
	client *rpc.Client
}

// Join the cluster through the http server
func (k *KioskClient) Join() {
	// TODO copy the necessary files e.g. /etc/corosync/corosync.conf
	// TODO run the pacemaker command to join the cluster once the cluster is configured

	// https://access.redhat.com/documentation/en-us/red_hat_enterprise_linux/7/html/high_availability_add-on_reference/s1-clusternodemanage-haar
	// TODO not hardcoded
	host := "node2"
	username := "hacluster"
	password := "admin"

	var reply string
	args := join.JoinArgs{
		Host:     host,
		Username: username,
		Password: password,
	}

	err := k.client.Call("Kiosk.Join", args, &reply)
	if err != nil {
		log.Fatalln("Error calling service Kiosk.Join", err)
	}
	log.Println(reply)
	pacemaker.StartCluser()
	pacemaker.EnableCluster()
	log.Println("The node", "node2", "has joined the cluster.")
	// TODO sync the Nix configuration

}

func Connect() *KioskClient {
	address := mdns_service.ReachCluster()
	client, err := rpc.DialHTTP("tcp", address)
	if err != nil {
		log.Fatalln("Cannot dial HTTP", err)
	}
	return &KioskClient{client: client}
}
