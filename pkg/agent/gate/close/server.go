package close

import (
	"log"

	"github.com/plmercereau/cluster-agent/pkg/agent/gate/sync"
	"github.com/plmercereau/cluster-agent/pkg/http_server"
	"github.com/plmercereau/cluster-agent/pkg/mdns_service"

	"github.com/plmercereau/cluster-agent/pkg/timer"
)

type CloseArgs struct{}

func CloseGate() {

	log.Println("Closing the gate...")
	timer.CancelTimer()
	// TODO stop the mDNS publishing too?
	http_server.StopHTTPServer()
	// wait for goroutine started in gate.StopServer() to stop
	sync.ServerExitDone.Wait()
	// httpServerExitDone.Done()
	mdns_service.Unpublish()

	log.Printf("Gate to join the cluster is now closed.")
}
