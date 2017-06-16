package parser

import (
	"fmt"
	"io/ioutil"

	"github.com/piokaczm/godzira/commands/task"
	"gopkg.in/yaml.v2"
)

const (
	rsync       = "rsync"
	rsyncArg    = "-chavzP"
	scp         = "scp"
	copyLabel   = "copy"
	localLabel  = "local"
	remoteLabel = "remote"
)

type config struct {
	Goos         string                    `yaml:"goos"`
	Goarch       string                    `yaml:"goarch"`
	Test         bool                      `yaml:"test"`
	Strategy     string                    `yaml:"strategy"`
	BinPath      string                    `yaml:"binary_path"`
	Environments map[string][]*environment `yaml:"environments"`
	PreTasks     []*unit                   `yaml:"pretasks"`
	PostTasks    []*unit                   `yaml:"posttasks"`
	env          string
}

func (c *config) interpretSingleTask(unit *unit) ([]*interpretedUnit, error) {
	var interpretedUnits []*interpretedUnit

	switch unit.label {
	case copyLabel:
		for _, host := range c.Environments[c.env] {
			interpretedUnit, err := unit.buildCopyCommand(host.Address(), c.Strategy)
			if err != nil {
				return nil, err
			}

			interpretedUnits = append(interpretedUnits, interpretedUnit)
		}
	case localLabel:
		interpretedUnits = append(interpretedUnits, unit.buildLocalCommand())
	case remoteLabel: // execute task for each host in the env
		for _, host := range c.Environments[c.env] {
			interpretedUnits = append(interpretedUnits, unit.buildRemoteCommand(host.Address()))
		}
	default:
		return nil, fmt.Errorf("[ command: %s ] '%s' label is not supported", unit.name, unit.label)
	}
	return interpretedUnits, nil
}

type environment struct {
	Host string `yaml:"host"`
	User string `yaml:"user"`
	Path string `yaml:"path"`
}

func (e *environment) Address() string {
	return fmt.Sprintf("%s@%s", e.User, e.Host)
}

func Read(queue *task.Queue, configPath, env string) error {
	// for config parts
	// - read a part of config
	// - check for special labels (local, remote, copy)
	// - interpret a single task
	// - create new task basing on interpretation
	// - append to a global queue

	conf, err := parse(configPath)
	if err != nil {
		return err
	}
	conf.env = env
	fmt.Println(conf)

	// r := &configReader{}

	return nil
}

func parse(configPath string) (*config, error) {
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

type configReader struct {
	errors []error
	queue  *task.Queue
}

func (cr *configReader) read(conf *config, env string) {
	cr.addTestTask(conf)

	for _, unit := range conf.PreTasks {
		interpretedUnits, err := conf.interpretSingleTask(unit)
		if err != nil {
			cr.errors = append(cr.errors, err)
			continue
		}

		for _, interpretedUnit := range interpretedUnits {
			cr.appendTask(interpretedUnit.name, interpretedUnit.command, task.PreTask)
		}
	}

	for _, unit := range conf.PostTasks {
		interpretedUnits, err := conf.interpretSingleTask(unit)
		if err != nil {
			cr.errors = append(cr.errors, err)
			continue
		}

		for _, interpretedUnit := range interpretedUnits {
			cr.appendTask(interpretedUnit.name, interpretedUnit.command, task.PostTask)
		}
	}

	err := cr.addDeployTask(conf)
	if err != nil {
		cr.errors = append(cr.errors, err)
	}
}

func (cr *configReader) addDeployTask(conf *config) error {
	for _, host := range conf.Environments[conf.env] {
		switch conf.Strategy {
		case rsync:
			cr.appendTask("deployment", fmt.Sprintf("%s %s %s %s", rsync, rsyncArg, conf.BinPath, host.Path), task.DeployTask)
		case scp:
			cr.appendTask("deployment", fmt.Sprintf("%s %s %s", scp, conf.BinPath, fmt.Sprintf("%s:%s", host.Address(), host.Path)), task.DeployTask)
		default:
			return unsupportedStrategy("deployment", conf.Strategy) // TODO: extract this deployment string ffs
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
	task, err := task.NewTask(name, command, taskType)
	if err != nil {
		cr.errors = append(cr.errors, err)
	}

	err = cr.queue.Append(task)
	if err != nil {
		cr.errors = append(cr.errors, err)
	}
}

func (cr *configReader) Fail() {
	// iterate over errors, print them and gracefully terminate the execution
}

func unsupportedStrategy(name, strategy string) error {
	return fmt.Errorf("[ command: %s ] '%s' strategy is not supported", name, strategy)
}
