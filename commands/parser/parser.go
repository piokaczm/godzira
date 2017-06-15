package parser

import (
	"fmt"
	"io/ioutil"

	"github.com/piokaczm/godzira/commands/task"
	"gopkg.in/yaml.v2"
)

type config struct {
	Goos         string                    `yaml:"goos"`
	Goarch       string                    `yaml:"goarch"`
	Test         bool                      `yaml:"test"`
	Strategy     string                    `yaml:"strategy"`
	BinName      string                    `yaml:"binary_name"`
	Environments map[string][]*environment `yaml:"environments"`
	PreTasks     map[string][]*task.Task   `yaml:"pretasks"`
	PostTasks    map[string][]*task.Task   `yaml:"posttasks"`
}

type environment struct {
	Host string `yaml:"host"`
	User string `yaml:"user"`
	Path string `yaml:"path"`
}

func Read(queue *task.Queue, configPath string) error {
	// for config parts
	// - read a part of config
	// - check for special labels (local, remote, copy)
	// - interpret a single task
	// - create new task basing on interpretation
	// - append to a global queue

	conf, err := readYaml(configPath)
	if err != nil {
		return err
	}
	fmt.Println(conf)

	return nil
}

func readYaml(configPath string) (*config, error) {
	conf := &config{}

	configData, err := ioutil.ReadFile(configPath)
	if err != nil {
		return conf, err
	}

	err = yaml.Unmarshal(configData, conf)
	if err != nil {
		return conf, err
	}

	return conf, nil
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
