package open

import (
	"log"
	"net"

	"sync"

	"github.com/plmercereau/cluster-agent/pkg/agent/gate/close"
	gateSync "github.com/plmercereau/cluster-agent/pkg/agent/gate/sync"
	"github.com/plmercereau/cluster-agent/pkg/kiosk_server"

	"github.com/plmercereau/cluster-agent/pkg/kiosk"
	"github.com/plmercereau/cluster-agent/pkg/mdns_service"
	"github.com/plmercereau/cluster-agent/pkg/timer"
)

type OpenArgs struct{}

// Open the cluster node so a new node can join
func OpenGate() {
	if kiosk_server.IsLive() {
		log.Println("Gate is already open. Keeping it open for", CLOSE_GATE_AFTER_SECONDS, "seconds more.")
		timer.PlanAction(close.CloseGate, CLOSE_GATE_AFTER_SECONDS)
		return
	}
	log.Println("Opening the gate to join the cluster...")
	listener := *kiosk_server.CreateListener()

	port := listener.Addr().(*net.TCPAddr).Port

	// TODO include mdns in the httpServerExitDone waitGroup too?
	go mdns_service.Publish(port)

	gateSync.ServerExitDone = &sync.WaitGroup{}

	gateSync.ServerExitDone.Add(1)
	kiosk.StartServer(listener, gateSync.ServerExitDone)

	// TODO concurrency: should wait for mDNS and HTTP server to be ready before returning

	log.Println("Gate is opened. It will automatically close in", CLOSE_GATE_AFTER_SECONDS, "seconds.")

	timer.PlanAction(close.CloseGate, CLOSE_GATE_AFTER_SECONDS)

}
