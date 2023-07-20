package http_server

import (
	"context"
	"log"
	"net"
	"net/http"
)

var server *http.Server

func IsLive() bool {
	return server != nil

}

func StartHttpServer(mux *http.ServeMux, listener net.Listener) error {
	server = &http.Server{
		Handler: mux,
	}

	return server.Serve(listener)

}

func StopHTTPServer() {
	log.Printf("Stopping HTTP server...")

	// now close the server gracefully ("shutdown")
	// timeout could be given with a proper context
	// TODO (in real world you shouldn't use TODO()).
	if err := server.Shutdown(context.TODO()); err != nil {
		panic(err) // failure/timeout shutting down the server gracefully
	}
	server = nil
	StopListener()
}
