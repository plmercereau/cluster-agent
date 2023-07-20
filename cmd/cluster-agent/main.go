package main

import (
	"log"
	"os"

	"github.com/plmercereau/cluster-agent/pkg/agent"
	"github.com/plmercereau/cluster-agent/pkg/agent/gate/close"
	"github.com/plmercereau/cluster-agent/pkg/agent/gate/open"
	"github.com/plmercereau/cluster-agent/pkg/kiosk"
)

func main() {

	switch os.Args[1] {
	case "start":
		agent.Start()

	case "open":
		open.OpenRequest()

	case "close":
		close.CloseRequest()

	case "join":
		client := kiosk.Connect()
		client.Join()

	default:
		log.Fatalln("Unknown command.")
	}

}
