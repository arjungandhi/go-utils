package shell

import "os/exec"

// CheckCommand checks if a command is in the PATH
func CheckCommand(command string) bool {
	_, err := exec.LookPath(command)
	if err != nil {
		return false
	}
	return true
}
