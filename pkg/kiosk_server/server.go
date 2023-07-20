package kiosk_server

import (
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
