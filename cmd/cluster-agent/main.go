package main

import (
	"github.com/plmercereau/cluster-agent/pkg/cluster"
	"github.com/plmercereau/cluster-agent/pkg/node"
	"github.com/plmercereau/cluster-agent/pkg/pacemaker"
)

// ? zeroconf / mDNS is probably overkill: we could just use multicast
func main() {
	if pacemaker.IsInsideCluster() {
		cluster.WatchNewNodes()
	} else {
		node.BroadcastNodeReady()
	}

}
