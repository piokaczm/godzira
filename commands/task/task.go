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

	preTask = iota
	deployTask
	postTask
)

// colors for printing messages
var (
	yellow = color.New(color.FgYellow).SprintFunc()
	red    = color.New(color.FgRed).SprintFunc()
)

// Task represents single step of deployment.
type Task struct {
	name     string   // name is a string which will be printed out during task execution
	cmd      *command // cms is a command struct which represents the actual command
	output   []byte   // output is cmd's output
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
		name: name,
		cmd:  cmd,
	}, err
}

func newCommand(cmd string) (*command, error) {
	cmdParts := strings.Split(cmd, " ")
	if len(cmdParts) < 1 { // prolly just check string len
		return nil, fmt.Errorf("malformed command -> %s", cmd)
	}

	return &command{
		name: cmdParts[0],
		args: cmdParts[1:],
	}, nil
}

// exec is a wrapper around task execution process.
func (t *Task) exec() error {
	t.print()
	return t.run()
}

// run is where a command is executed.
func (t *Task) run() error {
	output, err := exec.Command(t.cmd.name, t.cmd.args...).Output()
	t.output = output
	return err
}

func (t *Task) print() {
	fmt.Printf("%7s : executing task '%s'\n", yellow(time.Now().Format(format)), t.name)
}

func (t *Task) fail() {
	fmt.Printf("%7s : task failure '%s'\nCOMMAND OUTPUT: %s", yellow(time.Now().Format(format)), red(t.name), red(t.output))
}
