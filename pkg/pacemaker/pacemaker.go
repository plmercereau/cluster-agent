package pacemaker

import (
	"fmt"
	"log"
	"net"
	"os/exec"
)

func ActivateMaintenanceMode() error {
	log.Println("Activating maintenance mode...")
	cmd := exec.Command("pcs", "property", "set", "maintenance-mode=true")
	_, err := cmd.Output()
	return err

}

func DeactivateMaintenanceMode() error {
	log.Println("Deactivating maintenance mode...")
	cmd := exec.Command("pcs", "property", "set", "maintenance-mode=false")
	_, err := cmd.Output()
	return err

}

func AuthenticateNode(host string, address net.IP, username string, password string) error {
	// TODO use authkey instead of password. See: pcs cluster auth --help
	log.Println("Authenticating Node...")
	cmd := exec.Command("pcs", "host", "auth", host, fmt.Sprintf("addr=%s", address.String()), "-u", username, "-p", password)
	_, err := cmd.Output()
	return err
}

func AddNode(name string) error {
	log.Println("Adding node...")

	// * See the last part of https://clusterlabs.org/pacemaker/doc/deprecated/en-US/Pacemaker/2.0/html/Clusters_from_Scratch/_configure_corosync.html
	cmd := exec.Command("pcs", "cluster", "node", "add", name, "--start", "--enabe")
	_, err := cmd.Output()
	return err
}
