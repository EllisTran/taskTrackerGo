package command

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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
	description := getDescription(args)

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
	if len(args) > 1 {
		log.Fatalf("Too many commands: %d", len(args))
	}
	taskId := parseTaskId(args[0])
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

func updateFunc(cmd *Command, args []string) {
	taskId := parseTaskId(args[0])
	args = append(args[:0], args[1:]...)
	description := getDescription(args)
	updateTask(taskId, description)
}

func parseTaskId(id string) uint64 {
	const BitSize64 = 64
	const Base10 = 10
	taskId, err := strconv.ParseUint(id, Base10, BitSize64)
	if err != nil {
		log.Fatalf("Error: Task Id is not a number %v", err)
	}
	return taskId
}

// Parse string to see if it \u or not
func parseString(val string) string {

	if !strings.Contains(val, `\u`) {
		return val
	}
	fmt.Printf("in here")
	var decoded string
	err := json.Unmarshal([]byte(`"`+val+`"`), &decoded)
	if err != nil {
		return val
	}
	return decoded

}

func getDescription(words []string) string {
	description := ""
	for _, val := range words {

		parsedVal := parseString(val)
		description = description + " " + parsedVal
	}
	description = strings.TrimSpace(description)
	return description
}

func NewUpdateCommand() *Command {
	var updateUsage = `Update Task usage: brief update [OPTIONS]`
	cmd := &Command{
		Flags:   flag.NewFlagSet("update", flag.ExitOnError),
		Execute: updateFunc,
	}

	cmd.Flags.Usage = func() {
		fmt.Fprintln(os.Stderr, updateUsage)
	}
	return cmd

}

func listFunc(cmd *Command, args []string) {
	if len(args) == 0 {
		listTasks()
	} else if args[0] == "done" || args[0] == "in-progress" || args[0] == "todo" {
		listTasksWithStatus(TaskStatus(args[0]))
	}
}
func NewListCommand() *Command {
	var list = `List tasks usage: brief list [OPTIONS]`
	cmd := &Command{
		Flags:   flag.NewFlagSet("list", flag.ExitOnError),
		Execute: listFunc,
	}

	cmd.Flags.Usage = func() {
		fmt.Fprintln(os.Stderr, list)
	}
	return cmd
}

func markInProgressFunc(cmd *Command, args []string) {
	if len(args) > 1 {
		log.Fatalf("Too many commands: %d", len(args))
	}
	taskId := parseTaskId(args[0])
	markTask(taskId, TaskStatus(inProgress))
}

func NewMarkInProgressCommand() *Command {
	var markStatus = `Mark Status tasks usage: brief mark-in-progress [OPTIONS]`
	cmd := &Command{
		Flags:   flag.NewFlagSet("mark-in-progress", flag.ExitOnError),
		Execute: markInProgressFunc,
	}

	cmd.Flags.Usage = func() {
		fmt.Fprintln(os.Stderr, markStatus)
	}
	return cmd

}
func markDoneFunc(cmd *Command, args []string) {
	if len(args) > 1 {
		log.Fatalf("Too many commands: %d", len(args))
	}
	taskId := parseTaskId(args[0])
	markTask(taskId, TaskStatus(done))
}

func NewMarkDoneCommand() *Command {
	var markStatus = `Mark Status tasks usage: brief mark-done [OPTIONS]`
	cmd := &Command{
		Flags:   flag.NewFlagSet("mark-done", flag.ExitOnError),
		Execute: markDoneFunc,
	}

	cmd.Flags.Usage = func() {
		fmt.Fprintln(os.Stderr, markStatus)
	}
	return cmd

}
