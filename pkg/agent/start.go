package agent

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"os/signal"
	"syscall"

	"github.com/plmercereau/cluster-agent/pkg/agent/config"
	"github.com/plmercereau/cluster-agent/pkg/agent/gate"
)

func Start() {
	// Start the socket server so it awaits for instructions (open the cluster node to welcome new nodes)
	log.Println("Starting the agent...")
	os.RemoveAll(config.SOCKET_PATH)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		os.Remove(config.SOCKET_PATH)
		os.Exit(1)
	}()

	mux := http.NewServeMux()

	server := http.Server{
		Handler: mux,
	}

	rpcHandler := rpc.NewServer()
	welcomer := new(gate.Gate)
	rpcHandler.Register(welcomer)
	mux.Handle("/", rpcHandler)
	l, e := net.Listen("unix", config.SOCKET_PATH)
	if e != nil {
		log.Fatalln("listen error:", e)
	}

	log.Println("Agent started")

	server.Serve(l)

}
