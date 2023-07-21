package pacemaker

import "os/exec"

func IsInsideCluster() bool {
	// TODO not ideal. Also check if /etc/coronysync.conf exists?
	_, err := exec.Command("crm_mon", "-1").CombinedOutput()
	return err == nil
}
