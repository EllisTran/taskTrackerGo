package main

import (
	"flag"
	"fmt"
	"os"

	command "github.com/ellistran/taskTrackerGo/cmd"
	version "github.com/ellistran/taskTrackerGo/internal"
)

func usageAndExit(msg string) {
	if msg != "" {
		fmt.Fprint(os.Stderr, msg)
		fmt.Fprintf(os.Stderr, "\n")
	}

	flag.Usage()
	os.Exit(0)
}
func main() {
	var usage = `Usage: taskgo command [options] A simple tool to manage tasks. Options: Commands: add Adds a task. Delete delete a task in your tasks... etc etc`
	var cmd *command.Command
	switch os.Args[1] {
	case "version":
		cmd = version.NewVersionCommand()
	case "add":
		cmd = command.NewAddCommand()
	default:
		usageAndExit(fmt.Sprintf("taskgo: '%s' is not a taskgo command.\n", os.Args[1]))
	}
	cmd.Init(os.Args[2:])
	cmd.Run()
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprint(usage))
	}
	usageAndExit("")
}
