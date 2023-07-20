package mdns_service

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/grandcat/zeroconf"
	"github.com/plmercereau/cluster-agent/pkg/config"
)

func ReachCluster() string {
	log.Println("Looking for a cluster node with an open gate...")
	resolver, err := zeroconf.NewResolver(nil)
	if err != nil {
		log.Fatalln("Failed to initialize resolver:", err.Error())
	}

	entries := make(chan *zeroconf.ServiceEntry)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*config.TIMEOUT_IN_SECONDS)
	var address string
	go func(results <-chan *zeroconf.ServiceEntry) {
		for entry := range results {
			ips := entry.AddrIPv4
			ip := ips[len(ips)-1]
			// get the result of the goroutine
			address = fmt.Sprintf("%s:%d", ip, entry.Port)
			// Stop the resolver
			cancel()
		}

	}(entries)

	err = resolver.Lookup(ctx, config.INSTANCE_NAME, config.SERVICE_NAME, "local.", entries)
	defer cancel()
	if err != nil {
		log.Fatalln("Failed to browse:", err.Error())
	}

	<-ctx.Done()

	if address == "" {
		log.Fatalln("No cluster node found.")
	}
	log.Printf("Cluster node found: %s\n", address)
	return address
}
