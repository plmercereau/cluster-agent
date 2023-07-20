package kiosk

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"sync"

	"github.com/plmercereau/cluster-agent/pkg/kiosk_server"
)

func StartServer(l net.Listener, wg *sync.WaitGroup) {
	port := l.Addr().(*net.TCPAddr).Port
	log.Println("Starting HTTP server...")
	go func() {
		// defer l.Close()

		mux := http.NewServeMux()

		defer wg.Done() // let main know we are done cleaning up
		log.Println("Listening on port", port)

		joiner := new(Kiosk)
		handler := rpc.NewServer()
		handler.Register(joiner)
		mux.Handle("/", handler)

		err := kiosk_server.StartHttpServer(mux, l)
		log.Println("HTTP server stopped:", err)
	}()

}
