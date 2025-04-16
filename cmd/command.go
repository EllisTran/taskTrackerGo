package command

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
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

func deleteFunc(cmd *Command, args []string) {
	const BitSize64 = 64
	const Base10 = 10
	if len(args) > 1 {
		log.Fatalf("Too many commands: %d", len(args))
	}
	taskId, err := strconv.ParseUint(args[0], Base10, BitSize64)
	if err != nil {
		log.Fatalf("Error: Task Id is not a number %v", err)
	}
	fmt.Println("Delete task associated with id: %d", taskId)

	deleteTask(taskId)

}
func NewDeleteCommand() *Command {
	var deleteUsage = `Delete Task Usage: brief delete [OPTIONS] `
	cmd := &Command{
		Flags:   flag.NewFlagSet("delete", flag.ExitOnError),
		Execute: deleteFunc,
	}

	cmd.Flags.Usage = func() {
		fmt.Fprintln(os.Stderr, deleteUsage)
	}
	return cmd
}
