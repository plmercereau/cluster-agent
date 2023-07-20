package mdns_service

import (
	"log"

	"github.com/grandcat/zeroconf"
	"github.com/plmercereau/cluster-agent/pkg/config"
)

func Publish(port int) {
	// defer wg.Done()
	s, err := zeroconf.Register(config.INSTANCE_NAME, config.SERVICE_NAME, config.DOMAIN, port, []string{"txtv=0", "lo=1", "la=2"}, nil)
	if err != nil {
		log.Fatalln(err)
	}
	server = s
	log.Println("Published mDNS service at port", port)
	// defer server.Shutdown()

	// sig := make(chan os.Signal, 1)
	// signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	// <-sig
	// log.Println("Exit by user (mDNS)")
	// os.Exit(1)

}
