package parser

import (
	"fmt"
	"io/ioutil"

	"github.com/piokaczm/godzira/commands/task"
	"gopkg.in/yaml.v2"
)

const (
	rsync    = "rsync"
	rsyncArg = "-chavzP"
	scp      = "scp"
)

type config struct {
	Goos         string                    `yaml:"goos"`
	Goarch       string                    `yaml:"goarch"`
	Test         bool                      `yaml:"test"`
	Strategy     string                    `yaml:"strategy"`
	BinName      string                    `yaml:"binary_name"`
	Environments map[string][]*environment `yaml:"environments"`
	PreTasks     []*singleTask             `yaml:"pretasks"`
	PostTasks    []*singleTask             `yaml:"posttasks"`
}

type environment struct {
	Host string `yaml:"host"`
	User string `yaml:"user"`
	Path string `yaml:"path"`
}

type singleTask struct {
	name        string `yaml:"name"`
	command     string `yaml:"command"`
	path        string `yaml:"path"`
	destination string `yaml:"destination"`
	label       string `yaml:"type"`
}

func Read(queue *task.Queue, configPath, env string) error {
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

	r := &configReader{}

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
	queue  *task.Queue
}

func (cr *configReader) read(conf *config, env string) {
	// read a raw task and govern its' interpretation process
	// in case of an error just append it to errors which will be
	// printed out using Fail method.

	for _, task := range conf.PreTasks {
		err := cr.queue.Append(task.transposeToQueueTask())
		if err != nil {
			cr.errors = append(cr.errors, err)
		}
	}

	for _, task := range conf.PostTasks {
		err := cr.queue.Append(task.transposeToQueueTask())
		if err != nil {
			cr.errors = append(cr.errors, err)
		}
	}

	cr.addTestTask(config.Test)
	cr.addDeployTask(config)
}

func (cr *configReader) addDeployTask(conf *config) error {
	for _, environment := range c.Environments[env] {
		switch conf.Strategy {
		case rsync:
			cr.appendTask("deployment", fmt.Sprintf("%s %s", rsync, rsyncArg), task.DeployTask) // build command
		case scp:
			cr.appendTask("deployment", fmt.Sprintf("%s %s", scp, "?"), task.DeployTask) // build command
		default:
			return fmt.Errorf("%s startegy is not supported", conf.Strategy)
		}
	}

	return nil
}

func (cr *configReader) addTestTask(conf *config) {
	if conf.Test {
		cr.appendTask("run tests", "go test ./...", task.PreTask)
	}
}

func (cr *configReader) appendTask(name, command string, taskType int) {
	task := task.NewTask(name, command, taskType)
	err := cr.queue.Append(task)
	if err != nil {
		cr.errors = append(cr.errors, err)
	}
}

func (cr *configReader) Fail() {
	// iterate over errors, print them and gracefully terminate the execution
}
