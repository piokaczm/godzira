package main

import (
	"github.com/fatih/color"
	"os/exec"
)

type Task struct {
	Name         string
	Command      string
	Args         []string
	StartMessage string
	EndMessage   string
	Executed     bool
}

func getTask(name string, cmd string, args []string) *Task {
	return &Task{
		Name:         name,
		Command:      cmd,
		Args:         args,
		StartMessage: startMessage(name),
		EndMessage:   endMessage(name),
		Executed:     false,
	}
}

func (t *Task) exec() {
	t.printStart()
	t.runCommand()
	t.printEnd()
}

func (t *Task) runCommand() {
	output, err := exec.Command(t.Command, t.Args).Output()
	if err != nil {
		red := color.New(color.FgRed, color.Bold)
		red.Println(output)
		panic(err)
	}
	t.Executed = true
}

func (t *Task) printStart() {
	color.Yellow(t.EndMessage)
}

func (t *Task) printEnd() {
	green := color.New(color.FgGreen, color.Bold)
	green.Println(t.EndMessage)
}
