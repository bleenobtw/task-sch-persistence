package utils

import (
	"fmt"
	"os/exec"
	"strings"
)

func CheckIfTaskExists(taskName string) (bool, error) {
	cmd := exec.Command("schtasks", "/query", "/tn", taskName)
	output, err := cmd.CombinedOutput()

	if err != nil {
		if strings.Contains(string(output), "ERROR: The system cannot find the file specified.") {
			return false, nil
		}
		return false, fmt.Errorf("error checking Task Scheduler task existence: %v, output: %s", err, output)
	}
	return true, nil
}
