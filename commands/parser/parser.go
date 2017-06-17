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

func Read(conf *Config, queue *task.Queue) []error {
	r := &configReader{queue: queue}
	r.read(conf)
	return r.errors
}

func New(configPath, env string) (*Config, error) {
	conf := &Config{env: env}

	configData, err := ioutil.ReadFile(configPath)
	if err != nil {
		return conf, err
	}

	err = yaml.Unmarshal(configData, conf)
	if err != nil {
		return conf, fmt.Errorf("[ parsing ] an error occurred during parsing config file, please check if it's formatted correctly")
	}

	if conf.Environments[conf.env] == nil {
		return nil, fmt.Errorf("[ parsing ] '%s' couldn't find such an environment in configuration file", conf.env)
	}

	return conf, nil
}

type Config struct {
	Name         string                    `yaml:"name"`
	Goos         string                    `yaml:"goos"`
	Goarch       string                    `yaml:"goarch"`
	Test         bool                      `yaml:"test"`
	Strategy     string                    `yaml:"strategy"`
	BinPath      string                    `yaml:"binary_path"`
	Environments map[string][]*environment `yaml:"environments"`
	PreTasks     []*unit                   `yaml:"pretasks"`
	PostTasks    []*unit                   `yaml:"posttasks"`
	Slack        *Slack                    `yaml:"slack"`
	env          string
}

func (c *Config) interpretSingleTask(unit *unit) ([]*interpretedUnit, error) {
	var interpretedUnits []*interpretedUnit

	switch unit.Label {
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
		return nil, fmt.Errorf("[ command: %s ] '%s' label is not supported", unit.Name, unit.Label)
	}
	return interpretedUnits, nil
}

type Slack struct {
	Channel string `yaml:"channel"`
	Webhook string `yaml:"webhook"`
	Emoji   string `yaml:"emoji"`
	Name    string `yaml:"name"`
}

type environment struct {
	Host string `yaml:"host"`
	User string `yaml:"user"`
	Path string `yaml:"path"`
}

func (e *environment) Address() string {
	return fmt.Sprintf("%s@%s", e.User, e.Host)
}

type configReader struct {
	errors []error
	queue  *task.Queue
}

func (cr *configReader) read(conf *Config) {
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

func (cr *configReader) addDeployTask(conf *Config) error {
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

func (cr *configReader) addTestTask(conf *Config) {
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

func unsupportedStrategy(name, strategy string) error {
	return fmt.Errorf("[ command: %s ] '%s' strategy is not supported", name, strategy)
}
