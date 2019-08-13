package stat

import (
	"bytes"
	"os/exec"
)

// Runs a command, returns stdout if it executes successfully, stderr otherwise.
// Also returns an error which is nil for successful execution.
func runCmd(cmdName string, cmdArgs ...string) (string, error) {
	cmd := exec.Command(cmdName, cmdArgs...)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return string(stderr.Bytes()), err
	}

	return string(stdout.Bytes()), nil
}

func checkCmdExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	if err != nil {
		return false
	}
	return true
}
