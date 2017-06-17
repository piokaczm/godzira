package task

import (
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/fatih/color"
)

const (
	format = "2006-02-01 15:04:05"

	PreTask = iota
	DeployTask
	PostTask
)

// colors for printing messages
var (
	yellow     = color.New(color.FgYellow).SprintFunc()
	boldYellow = color.New(color.FgYellow).Add(color.Bold).SprintFunc()
	red        = color.New(color.FgRed).SprintFunc()
	boldRed    = color.New(color.FgRed).Add(color.Bold).SprintFunc()
	bold       = color.New(color.Bold).SprintFunc()
)

// Task represents single step of deployment.
type Task struct {
	name     string   // name is a string which will be printed out during task execution
	cmd      *command // cms is a command struct which represents the actual command
	output   []byte   // output is cmd's output
	err      error    // error stores errors raised during command execution for printing it in fail()
	taskType int      // taskType is used for assignment of a task to a proper queue
}

// command represents an actual command which will be executed during deployment process.
type command struct {
	name string   // name is an actual command, it's the first part of the command passed from config
	args []string // args are arguments for command
}

// NewTask is a Task constructor. If a command is malformed it will return an error.
func NewTask(name, command string, taskType int) (*Task, error) {
	cmd, err := newCommand(command)
	return &Task{
		name:     name,
		cmd:      cmd,
		taskType: taskType,
	}, err
}

func newCommand(cmd string) (*command, error) {
	cmdParts := strings.Split(cmd, " ")
	if len(cmd) < 1 {
		return nil, fmt.Errorf("malformed command : '%s'", cmd)
	}

	return &command{
		name: cmdParts[0],
		args: cmdParts[1:],
	}, nil
}

// exec is a wrapper around task execution process.
func (t *Task) exec() error {
	t.print()
	err := t.run()
	t.err = err
	return err
}

// run is where a command is executed.
func (t *Task) run() error {
	output, err := exec.Command(t.cmd.name, t.cmd.args...).Output()
	t.output = output
	return err
}

func (t *Task) print() {
	fmt.Printf("%7s : executing task '%s'\n", yellow(time.Now().Format(format)), bold(t.name))
}

func (t *Task) fail() {
	fmt.Printf(
		"%7s : task failure '%s'\nTASK OUTPUT:\n%s\n\nERROR:\n%s\n",
		yellow(time.Now().Format(format)),
		boldRed(t.name),
		red(string(t.output)),
		red(t.err),
	)
}
