package node

import (
	"crypto/rand"
	"encoding/base64"
	"log"
	"os"
	"sync"
	"time"

	"github.com/grandcat/zeroconf"
	"github.com/plmercereau/cluster-agent/pkg/config"
	"github.com/plmercereau/cluster-agent/pkg/pacemaker"
)

func generateRandomString(length int) string {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(b)
}

func BroadcastNodeReady() {
	log.Println("Letting know a potential cluster the node is ready to join...")
	wg := &sync.WaitGroup{}
	wg.Add(1)

	newPassword := generateRandomString(32)
	pacemaker.ChangeHAClusterPassword(newPassword)
	// ? encrypt password for the mDNS transport ?
	mdnsServer, errMDNS := zeroconf.Register(generateRandomString(10), config.SERVICE_NAME, config.DOMAIN, 1, []string{newPassword}, nil)
	defer mdnsServer.Shutdown()

	if errMDNS != nil {
		log.Fatalln("Failed to register a new mDNS service:", errMDNS.Error())
	}

	// No need to wg.Wait() as the following loop will run forever (until the node joins the cluster)
	for {
		if pacemaker.IsInsideCluster() {
			log.Println("Node joined the cluster.")
			os.Exit(0)
		}
		time.Sleep(1 * time.Second)
	}
}
