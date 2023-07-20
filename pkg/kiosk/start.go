package kiosk

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"sync"

	"github.com/plmercereau/cluster-agent/pkg/http_server"
)

func InitKiosk(l net.Listener, wg *sync.WaitGroup) {
	port := l.Addr().(*net.TCPAddr).Port
	log.Println("Starting HTTP server...")
	go func() {
		mux := http.NewServeMux()

		defer wg.Done() // let main know we are done cleaning up
		log.Println("Listening on port", port)

		joiner := new(Kiosk)
		handler := rpc.NewServer()
		handler.Register(joiner)
		mux.Handle("/", handler)

		http_server.StartHttpServer(mux, l)
		log.Println("HTTP server stopped.")
	}()

}
