package version

import (
	"fmt"
	bonzai "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)

var Cmd = &bonzai.Cmd{
	Name:     "name",
	Summary:  "summary",
	Commands: []*bonzai.Cmd{help.Cmd},
	Call: func(cmd *bonzai.Cmd, args ...string) error {
		fmt.Println(Version)
		return nil
	},
}
