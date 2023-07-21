package kiosk

import (
	"log"
	"net"
	"net/rpc"
	"os"

	"github.com/plmercereau/cluster-agent/pkg/kiosk/join"
	"github.com/plmercereau/cluster-agent/pkg/mdns_service"
)

type KioskClient struct {
	client *rpc.Client
}

// Join the cluster through the http server
func (k *KioskClient) Join() {
	// TODO copy the necessary files e.g. /etc/corosync/corosync.conf
	// TODO run the pacemaker command to join the cluster once the cluster is configured

	// https://access.redhat.com/documentation/en-us/red_hat_enterprise_linux/7/html/high_availability_add-on_reference/s1-clusternodemanage-haar
	hostname, _ := os.Hostname()

	var reply string
	args := join.JoinArgs{
		IP:   getOutboundIP(),
		Host: hostname,
		// TODO not hardcoded
		Username: "hacluster",
		Password: "admin",
	}

	err := k.client.Call("Kiosk.Join", args, &reply)
	if err != nil {
		log.Fatalln("Error calling service Kiosk.Join", err)
	}
	log.Printf("The node %s (%s) has joined the cluster.", args.Host, args.IP)
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

func getOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}
