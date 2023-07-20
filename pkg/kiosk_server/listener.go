package kiosk_server

import (
	"net"
)

var listener *net.Listener

func CreateListener() *net.Listener {
	address := ":0"
	if server != nil && listener != nil {
		address = (*listener).Addr().String()
		(*listener).Close()
	}
	l, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}
	listener = &l

	return listener
}

func StopListener() {
	if listener != nil {
		(*listener).Close()
	}
}
