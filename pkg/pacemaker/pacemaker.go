package pacemaker

import (
	"fmt"
	"log"
	"net"
	"os/exec"
	"strings"
)

func IsInsideCluster() bool {
	// TODO not ideal. Also check if /etc/coronysync.conf exists?
	_, err := exec.Command("crm_mon", "-1").CombinedOutput()
	return err == nil
}

func ChangeHAClusterPassword(password string) {
	// ! Make sure to add this line to /etc/sudoers:
	// ! user ALL=(ALL) NOPASSWD: /usr/sbin/usermod ^--password [^[:space:]]* hacluster$
	// ! Where user is the user running this program
	// TODO random salt
	sslOutput, err := exec.Command("openssl", "passwd", "-1", "-salt", "5RPVAd", password).Output()
	if err != nil {
		log.Fatalln("Error in encrypting the password", err, string(sslOutput))
	}
	encryptedPassword := strings.Replace(string(sslOutput), "\n", "", -1)

	output, err := exec.Command("sudo", "usermod", "--password", encryptedPassword, "hacluster").Output()
	if err != nil {
		log.Fatalln("Error in changing the password", err, string(output))
	}
	log.Println("Password changed.")
}

func ActivateMaintenanceMode() error {
	log.Println("Activating maintenance mode...")
	output, err := exec.Command("pcs", "property", "set", "maintenance-mode=true").CombinedOutput()
	if err != nil {
		log.Println("Error activating maintenance mode:", output)
	}
	return err

}

func DeactivateMaintenanceMode() error {
	log.Println("Deactivating maintenance mode...")
	output, err := exec.Command("pcs", "property", "set", "maintenance-mode=false").CombinedOutput()
	if err != nil {
		log.Println("Error deactivating maintenance mode:", output)
	}
	return err

}

func AuthenticateNode(host string, address net.IP, username string, password string) error {
	log.Printf("Authenticating %s (%s)...\n", host, address.String())
	output, err := exec.Command("pcs", "host", "auth", host, fmt.Sprintf("addr=%s", address.String()), "-u", username, "-p", password).CombinedOutput()
	if err != nil {
		log.Printf("Error authenticating %s: %s\n", host, output)
	}
	return err
}

func AddNode(name string) error {
	log.Printf("Adding %s...\n", name)
	// * See the last part of https://clusterlabs.org/pacemaker/doc/deprecated/en-US/Pacemaker/2.0/html/Clusters_from_Scratch/_configure_corosync.html
	output, err := exec.Command("pcs", "cluster", "node", "add", name, "--start", "--enable").CombinedOutput()
	if err != nil {
		log.Printf("Error adding %s: %s\n", name, output)
	}
	return err
}
