package parser

import (
	"fmt"

	"github.com/piokaczm/godzira/commands/task"
)

type unit struct {
	name        string `yaml:"name"`
	command     string `yaml:"command"`
	path        string `yaml:"path"`
	destination string `yaml:"destination"`
	label       string `yaml:"type"`
}

// local tasks doesn't have to be fired for each remote machine, so basing on provided
// labels godzira can decide if task should be appended
type interpretedUnit struct {
	name    string
	command string
}

func (u *unit) transposeToTask(taskType int, command string) (*task.Task, error) {
	return task.NewTask(u.name, command, taskType)
}

func (u *unit) buildCopyCommand(addr, strategy string) (interpretation *interpretedUnit, err error) {
	interpretation.name = u.name

	switch strategy {
	case rsync:
		interpretation.command = fmt.Sprintf("%s %s %s %s", rsync, rsyncArg, u.path, u.destination)
	case scp:
		interpretation.command = fmt.Sprintf("%s %s %s:%s", scp, u.path, addr, u.destination)
	default:
		err = unsupportedStrategy(u.name, strategy)
	}
	return
}

func (u *unit) buildLocalCommand() (interpretation *interpretedUnit) {
	interpretation.name = u.name
	interpretation.command = u.command
	return
}

func (u *unit) buildRemoteCommand(addr string) (interpretation *interpretedUnit) {
	interpretation.name = u.name
	interpretation.command = fmt.Sprintf("ssh %s %s", addr, u.command)
	return
}
