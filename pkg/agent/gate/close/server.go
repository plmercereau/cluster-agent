package close

import (
	"log"

	"github.com/plmercereau/cluster-agent/pkg/agent/gate/sync"
	"github.com/plmercereau/cluster-agent/pkg/kiosk_server"

	"github.com/plmercereau/cluster-agent/pkg/timer"
)

type CloseArgs struct{}

func CloseGate() {

	log.Println("Closing the gate...")
	// TODO stop the mDNS publishing too?
	kiosk_server.StopServer()
	// wait for goroutine started in gate.StopServer() to stop
	sync.ServerExitDone.Wait()
	// httpServerExitDone.Done()
	timer.CancelTimer()
	log.Printf("Gate to join the cluster is closed.")
}
