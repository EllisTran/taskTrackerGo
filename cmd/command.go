package command

import (
	"flag"
	"fmt"
	"os"
)

// Uppercase is private
// Lowercase is public
type Command struct {
	Flags   *flag.FlagSet
	Execute func(cmd *Command, args []string)
}

func (c *Command) Init(args []string) error {
	fmt.Println("Command: Init Method")
	return c.Flags.Parse(args)
}
func (c *Command) Called() bool {
	fmt.Println("Command: Called Method")
	return c.Flags.Parsed()
}
func (c *Command) Run() {
	fmt.Println("Command: Run Method")
	c.Execute(c, c.Flags.Args())
}

func addFunc(cmd *Command, args []string) {
	description := ""
	for _, val := range args {
		description = description + " " + val
	}

	fmt.Println("New Task: ", description)

	addTask(description, TaskStatus(todo))

}
func NewAddCommand() *Command {
	var addUsage = `Add a task Usage: brief add [OPTIONS] `
	cmd := &Command{
		Flags:   flag.NewFlagSet("add", flag.ExitOnError),
		Execute: addFunc,
	}

	cmd.Flags.Usage = func() {
		fmt.Fprintln(os.Stderr, addUsage)
	}
	return cmd
}
