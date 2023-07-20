package open

import (
	"log"
	"net"

	"sync"

	"github.com/plmercereau/cluster-agent/pkg/agent/gate/close"
	gateSync "github.com/plmercereau/cluster-agent/pkg/agent/gate/sync"
	"github.com/plmercereau/cluster-agent/pkg/config"
	"github.com/plmercereau/cluster-agent/pkg/http_server"

	"github.com/plmercereau/cluster-agent/pkg/kiosk"
	"github.com/plmercereau/cluster-agent/pkg/mdns_service"
	"github.com/plmercereau/cluster-agent/pkg/timer"
)

type OpenArgs struct{}

// Open the cluster node so a new node can join
func OpenGate() {
	if http_server.IsLive() {
		log.Println("Gate is already open. Keeping it open for", config.CLOSE_GATE_AFTER_SECONDS, "seconds more.")
		timer.PlanAction(close.CloseGate, config.CLOSE_GATE_AFTER_SECONDS)
		return
	}
	log.Println("Opening the gate to join the cluster...")
	l := *http_server.CreateListener()

	port := l.Addr().(*net.TCPAddr).Port

	// TODO include mdns in the httpServerExitDone waitGroup too?
	go mdns_service.Publish(port)

	gateSync.ServerExitDone = &sync.WaitGroup{}

	gateSync.ServerExitDone.Add(1)
	kiosk.InitKiosk(l, gateSync.ServerExitDone)

	// TODO concurrency: should wait for mDNS and HTTP server to be ready before returning

	log.Println("Gate is opened. It will automatically close in", config.CLOSE_GATE_AFTER_SECONDS, "seconds.")

	timer.PlanAction(close.CloseGate, config.CLOSE_GATE_AFTER_SECONDS)

}
