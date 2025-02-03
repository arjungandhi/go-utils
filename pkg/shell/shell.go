package shell

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

// CheckCommand checks if a command is in the PATH
func CheckCommand(command string) bool {
	_, err := exec.LookPath(command)
	if err != nil {
		return false
	}
	return true
}

func OpenInEditor(path string) {
	editor, exists := os.LookupEnv("EDITOR")
	if !exists {
		editor = "vi"
	}

	SysExec(editor, path)
}

// SysExec will check for the existence of the first argument as an
// executable on the system and then execute it using syscall.Exec(),
// which replaces the currently running program with the new one in all
// respects (stdin, stdout, stderr, process ID, signal handling, etc).
// Generally speaking, this is only available on UNIX variations.  This
// is exceptionally faster and cleaner than calling any of the os/exec
// variations, but it can make your code far be less compatible
// with different operating systems.
// source https://github.com/rwxrob/bonzai/blob/26f59ec373859d31411036b1208de8ac1e37782d/run/run.go#L123-L141
func SysExec(args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("missing name of executable")
	}
	path, err := exec.LookPath(args[0])
	if err != nil {
		return err
	}
	// exits the program unless there is an error
	return syscall.Exec(path, args, os.Environ())
}

// Exec checks for existence of first argument as an executable on the
// system and then runs it with [exec.Command.Run]  exiting in a way that
// is supported across all architectures that Go supports. The [os.Stdin],
// [os.Stdout], and [os.Stderr] are connected directly to that of the calling
// program. Sometimes this is insufficient and the UNIX-specific SysExec
// is preferred.
// source https://github.com/rwxrob/bonzai/blob/26f59ec373859d31411036b1208de8ac1e37782d/run/run.go#L143-L162
func Exec(args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("missing name of executable")
	}
	path, err := exec.LookPath(args[0])
	if err != nil {
		return err
	}
	cmd := exec.Command(path, args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
