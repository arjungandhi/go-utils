package version_test

import (
	"fmt"
	"testing"

	"github.com/arjungandhi/go-utils/pkg/version"
)

// Run these with go test -v
func TestVersion(t *testing.T) {
	fmt.Println(version.Version)
}
