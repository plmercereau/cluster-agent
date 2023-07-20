package pacemaker

import (
	"log"
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

func AuthenticateNode(host string, username string, password string) error {
	// TODO use authkey instead of password. See: pcs cluster auth --help
	log.Println("Authenticating Node...")
	cmd := exec.Command("pcs", "host", "auth", host, "-u", username, "-p", password)
	_, err := cmd.Output()
	return err
}

func AddNode(host string) error {
	log.Println("Adding node...")
	cmd := exec.Command("pcs", "cluster", "node", "add", host)
	_, err := cmd.Output()
	return err
}

func StartCluser() error {
	log.Println("Starting cluster...")
	cmd := exec.Command("pcs", "cluster", "start", "--all")
	_, err := cmd.Output()
	return err
}

func EnableCluster() error {
	log.Println("Enabling cluster...")
	cmd := exec.Command("pcs", "cluster", "enable", "--all")
	_, err := cmd.Output()
	return err
}
