package parser

import (
	"fmt"

	"github.com/piokaczm/godzira/commands/task"
	"gopkg.in/yaml.v2"
)

type config struct {
	goos         string                    `yaml:"goos"`
	goarch       string                    `yaml:"goarch"`
	test         bool                      `yaml:"test"`
	strategy     string                    `yaml:"strategy"`
	binName      string                    `yaml:"binary_name"`
	environments map[string][]*environment `yaml:"environments"`
	preTasks     map[string][]*task.Task   `yaml:"pretasks"`
	postTasks    map[string][]*task.Task   `yaml:"posttasks"`
}

type environment struct {
	host string `yaml:"host"`
	user string `yaml:"user"`
	path string `yaml:"path"`
}

func Read(queue *task.Queue) {
	// for config parts
	// - read a part of config
	// - check for special labels (local, remote, copy)
	// - interpret a single task
	// - create new task basing on interpretation
	// - append to a global queue

}

func parse() {
	// actually parse the config
	// return 3 data structures with raw data, one for each step of deployment
}

type configReader struct {
	errors []error
}

func (cr *configReader) read() {
	// read a raw task and govern its' interpretation process
	// in case of an error just append it to errors which will be
	// printed out using Fail method.
}

func (cr *configReader) Fail() {
	// iterate over errors, print them and gracefully terminate the execution
}
