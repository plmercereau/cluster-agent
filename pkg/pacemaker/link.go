package pacemaker

import (
	"log"
	"os/exec"
	"strings"
)

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
