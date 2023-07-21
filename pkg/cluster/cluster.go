package cluster

import (
	"context"
	"log"
	"strings"

	"github.com/grandcat/zeroconf"
	"github.com/plmercereau/cluster-agent/pkg/config"
	"github.com/plmercereau/cluster-agent/pkg/pacemaker"
)

func WatchNewNodes() {
	log.Println("Waiting for new nodes to join the cluster...")
	resolver, err := zeroconf.NewResolver(nil)
	if err != nil {
		log.Fatalln("Failed to initialize zeroconf resolver:", err.Error())
	}

	entries := make(chan *zeroconf.ServiceEntry)
	ctx := context.Background()

	go func(results <-chan *zeroconf.ServiceEntry) {
		for entry := range results {
			ips := entry.AddrIPv4
			// get the last local IP address
			// ? we may need to check every IP address
			ip := ips[len(ips)-1]
			hostname := strings.TrimSuffix(entry.HostName, "."+entry.Domain)
			password := entry.Text[0]
			maintenanceActivationError := pacemaker.ActivateMaintenanceMode()
			if maintenanceActivationError != nil {
				continue
			}
			authError := pacemaker.AuthenticateNode(hostname, ip, "hacluster", password)
			if authError != nil {
				continue
			}
			err := pacemaker.AddNode(hostname)
			// TODO run only if the maintenance mode was activated
			pacemaker.DeactivateMaintenanceMode()
			if err == nil {
				log.Println("Node added successfully.")
			}
		}

	}(entries)

	err = resolver.Browse(ctx, config.SERVICE_NAME, "local.", entries)

	if err != nil {
		log.Fatalln("Failed to browse mDNS services:", err.Error())
	}

	<-ctx.Done()
}
