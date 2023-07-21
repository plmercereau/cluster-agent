package main

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"

	"github.com/grandcat/zeroconf"
	"github.com/plmercereau/cluster-agent/pkg/config"
	"github.com/plmercereau/cluster-agent/pkg/pacemaker"
)

func IsInCluster() bool {
	// TODO
	isInCluster := os.Getenv("LOCATION") == "cluster"
	if isInCluster {
		log.Println("In the cluster.")
	} else {
		log.Println("Outside of the cluster.")

	}
	return isInCluster
}

func ChangeHAClusterPassword(password string) {
	// ! Make sure to add this line to /etc/sudoers:
	// ! user ALL=(ALL) NOPASSWD: /usr/sbin/usermod ^--password [^[:space:]]* hacluster$
	// ! Where user is the user running this program
	// TODO random salt
	sslCmd := exec.Command("openssl", "passwd", "-1", "-salt", "5RPVAd", password)
	sslOutput, err := sslCmd.Output()
	if err != nil {
		log.Fatalln("Error in encrypting the password", err, string(sslOutput))
	}
	encryptedPassword := strings.Replace(string(sslOutput), "\n", "", -1)

	cmd := exec.Command("sudo", "usermod", "--password", encryptedPassword, "hacluster")
	output, err := cmd.Output()
	if err != nil {
		log.Fatalln("Error in changing the password", err, string(output))
	}
	log.Println("Password changed.")
}

func generateRandomString(length int) string {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(b)
}

// ? zeroconf / mDNS is probably overkill: we could just use multicast
// TODO do not add the node that way (mdns) but through an http api, as we must shut down the service in the node after it joined.
func main() {
	if IsInCluster() {
		log.Println("Waiting for new nodes to join the cluster...")
		resolver, err := zeroconf.NewResolver(nil)
		if err != nil {
			log.Fatalln("Failed to initialize zeroconf resolver:", err.Error())
		}

		entries := make(chan *zeroconf.ServiceEntry)
		ctx := context.Background()

		go func(results <-chan *zeroconf.ServiceEntry) {
			for entry := range results {
				ips := entry.AddrIPv4
				ip := ips[len(ips)-1]
				hostname := strings.TrimSuffix(entry.HostName, "."+entry.Domain)
				password := entry.Text[0]
				maintenanceActivationError := pacemaker.ActivateMaintenanceMode()
				if maintenanceActivationError != nil {
					continue
				}
				authError := pacemaker.AuthenticateNode(hostname, ip, "hacluster", password)
				if authError != nil {
					continue
				}
				err := pacemaker.AddNode(hostname)
				if err == nil {
					log.Println("Node added successfully.")
				}
				// TODO run only if the maintenance mode was activated
				pacemaker.DeactivateMaintenanceMode()
			}

		}(entries)

		err = resolver.Browse(ctx, config.SERVICE_NAME, "local.", entries)

		if err != nil {
			log.Fatalln("Failed to browse:", err.Error())
		}

		<-ctx.Done()

	} else {

		wg := &sync.WaitGroup{}
		wg.Add(1)

		newPassword := generateRandomString(32)
		ChangeHAClusterPassword(newPassword)
		// TODO encrypt password for the transport
		mdnsServer, errMDNS := zeroconf.Register(generateRandomString(10), config.SERVICE_NAME, config.DOMAIN, 1, []string{newPassword}, nil)
		defer mdnsServer.Shutdown()

		if errMDNS != nil {
			log.Fatalln(errMDNS)
		}

		wg.Wait()
	}

}
