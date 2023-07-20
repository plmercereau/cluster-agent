package kiosk_server

import (
	"context"
	"log"

	"github.com/plmercereau/cluster-agent/pkg/mdns_service"
)

func StopServer() {
	log.Printf("Stopping HTTP server...")

	// now close the server gracefully ("shutdown")
	// timeout could be given with a proper context
	// (in real world you shouldn't use TODO()).
	if err := server.Shutdown(context.TODO()); err != nil {
		panic(err) // failure/timeout shutting down the server gracefully
	}
	server = nil
	StopListener()
	mdns_service.Unpublish()

}
