package version

import (
	"flag"
	"fmt"
	"os"

	command "github.com/ellistran/taskTrackerGo/cmd"
)

var versionUsage = `Print the app version and build info for the current context. Usage: taskgo version [options] Options: --short If true, print just the version number. Default false. `

// What might've caught your eye is the ??? for the build and version variables. Those are meant as placeholder variables that will be overwritten during the build process. This allows us to specify version numbers in a config or dynamically include commit hashes on each build.
var (
	build   = "???"
	version = "???"
	short   = false
)

var versionFunc = func(cmd *command.Command, args []string) {
	if short {
		fmt.Printf("brief version: v%s", version)
	} else {
		fmt.Printf("brief version: v%s, build: %s", version, build)
	}
	os.Exit(0)
}

func NewVersionCommand() *command.Command {

	cmd := &command.Command{
		Flags:   flag.NewFlagSet("version", flag.ExitOnError),
		Execute: versionFunc,
	}

	cmd.Flags.BoolVar(&short, "short", false, "Print just the version number")
	cmd.Flags.Usage = func() {
		fmt.Fprintln(os.Stderr, versionUsage)
	}

	return cmd
}
