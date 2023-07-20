package mdns_service

import (
	"log"
)

func Unpublish() {
	server.Shutdown()

	log.Println("Published mDNS service is shut down.")

}
