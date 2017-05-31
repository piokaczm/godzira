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

// Task represents single step of deployment
type Task struct {
	name       string
	cmd        *command
	output     []byte
	taskType   int
	successful bool
}

type command struct {
	name string
	args string
}

func NewTask(name, command string, taskType int) *Task {
	cmd, err := newCommand(command)
	return &Task{
		name:    name,
		command: cmd,
		args:    args,
	}
}

func newCommand(cmd string) (*command, error) {
	cmdParts := strings.Split(command, " ")
	if len(cmdParts) < 1 { // prolly just check string len
		return nil, fmt.Errorf("malformed command")
	}

	return &command{
		name: cmdParts[0],
		args: cmdParts[1:],
	}, nil
}

// Exec is a wrapper around task execution process.
func (t *Task) Exec() {
	t.print()
	t.run()

	if !t.successful {
		t.fail()
	}
}

func (t *Task) run() {
	output, err := exec.Command(t.cmd.name, t.cmd.args...).Output()
	t.output = output

	if err == nil {
		t.successful = true
	}
}

func (t *Task) print() {
	fmt.Printf("%7s : executing task '%s'\n", yellow(time.Now().Format(format)), t.name)
}

func (t *Task) fail() {
	fmt.Printf("%7s : task failure '%s'\nCOMMAND OUTPUT: %s", yellow(time.Now().Format(format)), red(t.name), red(t.output))
}
