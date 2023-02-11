package shell_test

import (
	"testing"

	"github.com/arjungandhi/go-utils/pkg/shell"
)

func TestShell(t *testing.T) {
	if !shell.CheckCommand("ls") {
		t.Error("ls is not in the PATH")
	}

	if shell.CheckCommand("not-a-command") {
		t.Error("not-a-command is in the PATH")
	}
}
